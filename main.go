package main

import (
	"context"

	usecases "for_learning/src/app/usecase"
	bookUC "for_learning/src/app/usecase/books"
	"for_learning/src/infra/config"
	"for_learning/src/infra/persistence/redis"

	"for_learning/src/interface/rest"

	bookInteg "for_learning/src/infra/integration/books"

	ms_log "for_learning/src/infra/log"

	circuit_breaker_service "for_learning/src/infra/circuit_breaker"
	redisService "for_learning/src/infra/persistence/redis/service"

	_ "github.com/joho/godotenv/autoload"
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

	circuitBreaker := circuit_breaker_service.NewCircuitBreakerInstance()
	bookIntegration := bookInteg.NewIntegOpenLibrary(circuitBreaker)

	// HTTP Handler
	// the server already implements a graceful shutdown.

	allUC := usecases.AllUseCases{
		BookUC: bookUC.NewBooksUseCase(bookIntegration, redisSvc),
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
