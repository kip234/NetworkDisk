package Redis

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
	rpool *redis.Pool//负责读取
	wpool *redis.Pool//负责写入
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
	if v==nil {
		return "",nil
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

func (r RedisPool)SADD(args ...interface{}) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("SADD",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	return nil
}

func (r RedisPool)SISMEMBER(args ...interface{}) (int64,error) {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("SISMEMBER",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	re,ok:=v.(int64)
	if !ok {
		return 0,fmt.Errorf("func (r RedisPool)SISMEMBER(args ...interface{}) (int,error) : Assertion error")
	}
	return re,nil
}

func (r RedisPool)SMEMBERS(key string) (re []string,err error) {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("SMEMBERS",key)
	fmt.Println(v)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)SMEMBERS(key string) ([]string,error): %s",err.Error())
	}
	tmp:=v.([]interface{})
	for _,i:=range tmp{
		re=append(re, string(i.([]uint8)))//这里很有可能会出问题
	}
	return re,nil
}