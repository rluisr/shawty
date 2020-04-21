package models

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	RedisAddr               string
	RedisSentinelMasterName string
	RedisSentinelAddr       []string
	RedisPassword           string
	RedisDB                 int
	GenerateSize            int
}

func NewConfig() *Config {
	redisDBInt, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}

	generateSize, err := strconv.Atoi(os.Getenv("GENERATE_SIZE"))
	if err != nil {
		panic(err)
	}

	var redisSentinelAddr []string
	if os.Getenv("REDIS_SENTINEL_ADDR") != "" {
		redisSentinelAddr = strings.Split(os.Getenv("REDIS_SENTINEL_ADDR"), ",")
	}

	return &Config{
		RedisAddr:               os.Getenv("REDIS_ADDR"),
		RedisSentinelMasterName: os.Getenv("REDIS_SENTINEL_MASTER_NAME"),
		RedisSentinelAddr:       redisSentinelAddr,
		RedisPassword:           os.Getenv("REDIS_PASSWORD"),
		RedisDB:                 redisDBInt,
		GenerateSize:            generateSize,
	}
}
