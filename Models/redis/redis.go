package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisPool struct {
	Read 		string
	Write 		string
	IdLeTimeout	int
	MaxIdle		int
	MaxActive	int
	rpool *redis.Pool
	wpool *redis.Pool
}

func (r *RedisPool)Init()  {
	r.rpool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn,err := redis.Dial("tcp",r.Read)
			return conn,err
		},
		MaxIdle: r.MaxIdle,
		MaxActive: r.MaxActive,
		IdleTimeout: time.Second*time.Duration(r.IdLeTimeout),
	}

	r.wpool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn,err := redis.Dial("tcp",r.Write)
			return conn,err
		},
		MaxIdle: r.MaxIdle,
		MaxActive: r.MaxActive,
		IdleTimeout: time.Second*time.Duration(r.IdLeTimeout),
	}
}

func (r RedisPool)SET(args ...interface{}) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("SET",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)SET(args ...interface{}) error: %s",err.Error())
	}
	return err
}

func (r RedisPool)GET(key string) (string,error) {
	rdb :=r.rpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("GET",key)
	if err != nil{
		return "",err
	}
	re,ok:=v.([]uint8)
	if !ok {
		return "",fmt.Errorf("func (r RedisPool)GET(key string) (string,error) : Assertion failure")
	}
	return string(re),err
}

func (r RedisPool)DEL(key string) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("DEL",key)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	return err
}