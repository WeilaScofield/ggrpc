package redis

import (
	"github.com/go-redis/redis"
	"log"
)

func NewRedisClient(addr string, password string, db int) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("redis connection error")
	}
	return
}