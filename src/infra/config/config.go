package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port string
	
	Timeout int
}

type LogConf struct {
	Name string
}

type RPSConf struct {
	Limit int
}

// Config ...
type Config struct {
	App  AppConf
	Http HttpConf
	Log  LogConf
	RPS  RPSConf
}

// NewConfig ...
func Make() Config {
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	http := HttpConf{
		Port: os.Getenv("HTTP_PORT"),
	}

	log := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	limit, _ := strconv.Atoi(os.Getenv("MAX_REQUEST_LIMIT"))

	rps := RPSConf{
		Limit: limit,
	}
	config := Config{
		App:  app,
		Http: http,
		Log:  log,
		RPS:  rps,
	}

	return config
}
