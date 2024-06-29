package main

import (
	"context"
	"log"
	"time"

	usecases "for_learning/src/app/usecase"
	bookUC "for_learning/src/app/usecase/books"
	bookGrpcUC "for_learning/src/app/usecase/books_grpc"
	pickUpUC "for_learning/src/app/usecase/pickup"
	"for_learning/src/infra/config"
	"for_learning/src/infra/persistence/redis"

	"for_learning/src/interface/rest"

	booksProto "for_learning/src/app/proto/books"
	bookInteg "for_learning/src/infra/integration/books"
	bookIntegGrpc "for_learning/src/infra/integration/books_grpc"

	ms_log "for_learning/src/infra/log"

	circuit_breaker_service "for_learning/src/infra/circuit_breaker"
	redisService "for_learning/src/infra/persistence/redis/service"

	"for_learning/src/infra/broker/nats"
	natsPublisher "for_learning/src/infra/broker/nats/publisher"

	_ "github.com/joho/godotenv/autoload"
	grpc "google.golang.org/grpc"
)

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	redisClient, err := redis.NewRedisClient(conf.Redis, logger)

	redisSvc := redisService.NewServRedis(redisClient)

	// Buat koneksi gRPC
	gRPCConn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer gRPCConn.Close()

	// Buat client dari BookService
	gRPCClient := booksProto.NewBookServiceClient(gRPCConn)
	circuitBreaker := circuit_breaker_service.NewCircuitBreakerInstance()
	bookIntegration := bookInteg.NewIntegOpenLibrary(circuitBreaker)
	bookIntegrationGrpc := bookIntegGrpc.NewIntegOpenLibrary(circuitBreaker, gRPCClient)

	Nats := nats.NewNats(conf.Nats, logger)
	publisher := natsPublisher.NewPushWorker(Nats)
	// HTTP Handler
	// the server already implements a graceful shutdown.

	allUC := usecases.AllUseCases{
		BookUC:     bookUC.NewBooksUseCase(bookIntegration, redisSvc),
		PickUpUC:   pickUpUC.NewPickUpUseCase(publisher),
		BookGrpcUC: bookGrpcUC.NewBooksGRPCUseCase(bookIntegrationGrpc, redisSvc),
	}

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		allUC,
		conf.RPS,
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
