package main

import (
	"fmt"
	"go.uber.org/zap"
	"github.com/labstack/echo/v4"
	"github.com/vaibhavqwerty/mini-redis/cmd/handlers"
)

func main(){
	var x int32 = 0
	fmt.Println("hello")
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Fatal("couldn't flush zap logger", zap.Error(err))
		}
	}()

	e := echo.New()

	checkHealthHandler := handlers.NewHealth(logger,&x)
	e.GET("/healthz",checkHealthHandler.Handle)

	if err := e.Start(fmt.Sprintf(":%d", 8081)); err != nil {
		logger.Fatal("---Error While Serving---", zap.Error(err))
	}
}