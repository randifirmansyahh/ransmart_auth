package redisHelper

import (
	"context"
	"encoding/json"
	"log"
	"ransmart_auth/app/helper/helper"
	"ransmart_auth/app/models/userModel"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	ctx            = context.Background()
	redissetfailed = "Gagal set data ke redis"
)

// del redis
func ClearRedis(REDIS *redis.Client, key string) {
	REDIS.Del(ctx, key)
}

// set ke redis
func SetRedis(REDIS *redis.Client, keyRedis string, result []byte) {
	if err := REDIS.Set(ctx, keyRedis, (result), 0).Err(); err != nil {
		log.Println(redissetfailed)
		return
	}
}

// get redis data by key
func getRedis(REDIS *redis.Client, key string) (string, error) {
	response, err := REDIS.Get(ctx, key).Result()
	return response, err
}

// get redis data with response
func GetRedisData(key_redis string, redis *redis.Client) ([]interface{}, error) {
	// get redis
	result, err := getRedis(redis, key_redis)
	if err != nil {
		return nil, err
	}

	// assign to model
	var data []interface{}
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		return nil, err
	}

	// set to redis
	if newMarshall, err := json.Marshal(data); err == nil {
		SetRedis(redis, key_redis, newMarshall)
	}

	// response.ResponseSuccess(w, handlerName, data)
	return data, nil
}

func GetOneRedisData(id, key_redis string, redis *redis.Client) (interface{}, error) {
	// search from redis
	data, err := getRedis(redis, key_redis)
	if err != nil {
		return nil, err
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// unmarshall
	var result []userModel.User
	json.Unmarshal([]byte(data), &result)

	// search from data
	if oneData, err := helper.SearchOneUser(result, newId); err != nil {
		return nil, err
	} else {
		return oneData, nil
	}
}
