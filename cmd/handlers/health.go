package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"fmt"
)

type HealthStatus struct {
	log *zap.Logger
	x *int32
}

func NewHealth(log *zap.Logger, x *int32) HealthStatus{
	return HealthStatus{
		log: log,
		x: x,
	}
}

func (h HealthStatus) Handle(c echo.Context) (err error){
	
	if 1==1{

		h.log.Info("---health check---")
		*h.x=*(h.x)+1
		fmt.Println(*h.x)
		return c.String(http.StatusOK,"OK")
	}

	return  c.String(http.StatusInternalServerError, "Mini-redis is unhealthy")
}