package main

import (
	"fmt"
	"go.uber.org/zap"
	"github.com/labstack/echo/v4"
	"github.com/vaibhavqwerty/mini-redis/cmd/handlers"
	"github.com/vaibhavqwerty/mini-redis/internal/api"
)

func main(){
	fmt.Println("hello")
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Fatal("couldn't flush zap logger", zap.Error(err))
		}
	}()

	e := echo.New()

	redisObject := api.NewRedisObj()

	checkHealthHandler := handlers.NewHealth(logger)
	e.GET("/healthz",checkHealthHandler.Handle)

	redisHandler := handlers.NewRedis(&redisObject,logger)
	e.POST("/redis",redisHandler.Handle)

	if err := e.Start(fmt.Sprintf(":%d", 8083)); err != nil {
		logger.Fatal("---Error While Serving---", zap.Error(err))
	}
}