package api

import (
	"github.com/vaibhavqwerty/mini-redis/internal/model"
)

type RedisObj struct{
	valStore map[string]string
	timeStore map[string]model.TimeStore
}

func NewRedisObj() RedisObj{
	return RedisObj{
		valStore: make(map[string]string),
		timeStore: make(map[string]model.TimeStore),
	}
}


func (r RedisObj)get(key string) string{
	return r.valStore[key]
}

func (r RedisObj)set(key string, val string) string{
	r.valStore[key]=val
	return key
}