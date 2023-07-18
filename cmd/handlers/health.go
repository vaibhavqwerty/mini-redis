package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type HealthStatus struct {
	log *zap.Logger
}

func NewHealth(log *zap.Logger) HealthStatus{
	return HealthStatus{
		log: log,
	}
}

func (h HealthStatus) Handle(c echo.Context) (err error){
	
	if true{

		h.log.Info("----health check---")
		return c.String(http.StatusOK,"OK")
	}

	return  c.String(http.StatusInternalServerError, "Mini-redis is unhealthy")
}