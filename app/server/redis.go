package server

import (
	"context"
	"log"
	"os"
	"ransmart_auth/app/helper/helper"

	"github.com/go-redis/redis/v8"
)

func connectToRedis() *redis.Client {
	// insialisasi Koneksi
	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),                   // hostname
		Password: os.Getenv("REDIS_PASSWORD"),               // password
		DB:       helper.StringToint(os.Getenv("REDIS_NO")), // bisa 0,1,2,3,4,5,6,7,8,9 dll
	})

	ctx := context.Background()

	// cek redis asynchronous
	msg, err := redis.Ping(ctx).Result() // Test koneksi Redis nya (nyala atau engga)
	if err != nil || msg != "PONG" {
		log.Println("not conect error =>", err)
		log.Println("Redis Not Connected !!")
	} else {
		log.Println("Redis Connected")
	}

	return redis
}
