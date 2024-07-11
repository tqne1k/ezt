package main

import (
	"eztrust/api/route"
	"eztrust/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	env := bootstrap.NewEnv()
	bootstrap.InitDabaBase(env)
	database := bootstrap.Connect(env)
	bootstrap.Migrate(database)
	timeout := time.Duration(env.ContextTimeout) * time.Second
	loggerConfig := bootstrap.LogOptions{
		ConsoleLoggingEnabled: true,
		FileLoggingEnabled:    true,
		Directory:             "log",
		Filename:              "eztrust.log",
		MaxSize:               100,
		MaxBackups:            120,
		MaxAge:                120,
	}
	bootstrap.ConfigureLogger(loggerConfig)

	//	web api
	gin := gin.Default()
	route.Setup(env, timeout, database, gin)
	bootstrap.Logger.Info().Msg("Starting API  server...")
	gin.Run(env.ServerAddress)
}
