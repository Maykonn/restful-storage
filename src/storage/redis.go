package storage

import (
	"github.com/go-redis/redis"
	"os"
	"sync"
)

var RedisCli *redis.Client
var once sync.Once

// Singleton for Redis.Client
func Redis() *redis.Client {
	// thread-safe
	once.Do(func() {
		if address := os.Getenv("REDIS_ADDRESS"); address != "" {
			RedisCli = redis.NewClient(&redis.Options{
				Addr:     address,
				Password: "",
				DB:       0,
			})
		} else {
			panic("Invalid Redis Server configuration")
		}
	})
	return RedisCli
}
