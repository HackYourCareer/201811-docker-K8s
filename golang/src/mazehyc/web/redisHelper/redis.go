package redisHelper

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func NewClient(addr string) *redis.Client {
	if addr == "" {
		panic("Redis connection address missing")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("CONNECTED")
	return client
}

func GetValue(client *redis.Client, key string) (bool, string, error) {
	res, err := client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			//Not Found!
			return false, "", nil
		}

		return false, "", err
	}

	return true, res, nil
}

func SetValue(client *redis.Client, key, value string, ttlSeconds uint) error {
	err := client.Set(key, value, time.Second*time.Duration(ttlSeconds)).Err()
	if err != nil {
		return err
	}
	return nil
}

type GetRedisValueFunc func(key string) (bool, string, error)
type SetRedisValueFunc func(key, value string) error
