package api

import (
	"github.com/vaibhavqwerty/mini-redis/internal/model"
	"regexp"
	"time"
	"strconv"
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


func (r *RedisObj)Get(key string) string{
	val,ok1 := r.timeStore[key]
	if ok1{
		diff := time.Since(val.EntryTime)
		dur2 := diff.Seconds()
		if(dur2>=float64(val.Duration)){
			delete(r.valStore,key)
			delete(r.timeStore,key)
			return "-2"
		}
	}

	return r.valStore[key]
	
}

func (r *RedisObj)Set(key string, val string) string{
	r.valStore[key]=val
	return "OK"
}

func (r *RedisObj)Del(key string) string{
	_,ok := r.valStore[key]
	_,ok1 := r.timeStore[key]
	if ok{
		delete(r.valStore,key)
		if ok1{
			delete(r.timeStore,key)	
		}
		return "1"
	}

	return "0"	
}

func(r *RedisObj)Keys(st string) string{
	regex, _ := regexp.Compile(st)
	ans := ""
	for key := range r.valStore{

		val,ok1 := r.timeStore[key]

		if ok1{
			diff := time.Since(val.EntryTime)
			dur2 := diff.Seconds()
			if(dur2>=float64(val.Duration)){
				delete(r.valStore,key)
				delete(r.timeStore,key)
				continue
			}
		}

		if regex.MatchString(key){
			if ans==""{
				ans=ans+key;
			}else{
				ans+=" , "+key;
			}
		}
	}

	return ans;
}

func(r *RedisObj)Expires(key string, duration int64) string{
	_,ok := r.valStore[key]

	if !ok{
		return "0"
	}else{
		x := model.TimeStore{
			EntryTime: time.Now().UTC(),
			Duration: duration,
		}
		r.timeStore[key]= x
	}

	return "1"

}

func (r *RedisObj)Ttl(key string) string{
	_,ok := r.valStore[key]
	val,ok1 := r.timeStore[key]

	if !ok{
		return "-2"
	}
	if ok && !ok1{
		return "-1"
	}
	if ok1{
		diff := time.Since(val.EntryTime)
		dur2 := diff.Seconds()
		if(dur2>=float64(val.Duration)){
			delete(r.valStore,key)
			delete(r.timeStore,key)
			return "-2"
		}else{
			return strconv.FormatFloat(float64(val.Duration)-dur2,'f', -1, 64)
		}

	}
	
	return ""

}