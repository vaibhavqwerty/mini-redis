package handlers

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/vaibhavqwerty/mini-redis/internal/api"
	"go.uber.org/zap"
	"regexp"
	"strconv"
)

type Redis struct {
	svc *api.RedisObj
	log *zap.Logger
}

func NewRedis(svc *api.RedisObj, log *zap.Logger) Redis {
	return Redis{
		svc: svc,
		log: log,
	}
}

func (r Redis) Handle(c echo.Context) (err error) {

	query := c.FormValue("query")

	if query == "" {
		// Return a custom error response
		r.log.Info("--- no query provided ---")
		return c.String(http.StatusBadRequest, "error (empty query)")
	}

	words := strings.Fields(query)

	baseCmd := words[0]

	result := "nil"
	isCorrectQuery := true
	switch baseCmd {
	case "GET":
		if len(words) == 2 {
			result = r.svc.Get(words[1])
		} else {
			isCorrectQuery = false
		}

	case "SET":
		if len(words) == 3 {
			result = r.svc.Set(words[1], words[2])
		} else {
			isCorrectQuery = false
		}
	case "DEL":
		if len(words)==2 {
			result = r.svc.Del(words[1])
		} else{
			isCorrectQuery = false
		}
	case "KEYS":
		if(len(words)==2){
			pattern := words[1]
			_, er := regexp.Compile(pattern)
			if er != nil {
				fmt.Println(er)
				r.log.Info("---Error compiling regex pattern---")
				isCorrectQuery = false
			}else{
				r.svc.Keys(words[1])
			}
		}else{
			isCorrectQuery = false
		}
	case "EXPIRES":
		if(len(words)==3){
			num, er := strconv.ParseInt(words[2], 10, 64)
			if er != nil {
				isCorrectQuery = false
			}else{
				result=r.svc.Expires(words[1],num)
			}
			
		}else{
			isCorrectQuery = false
		}
	case "TTL":
		if(len(words)==2){
			result=r.svc.Ttl(words[1])
		}else{
			isCorrectQuery = false
		}
	default:
		r.log.Info("--- error in query ---")
		isCorrectQuery=false
	}

	if !isCorrectQuery {
		return c.String(http.StatusBadRequest, "error (invalid query)")

	}
	if result == "" {
		result = "(nil)"
	}
	return c.String(http.StatusOK, result)
}
