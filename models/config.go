package models

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	GenerateSize  int
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

	return &Config{
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDBInt,
		GenerateSize:  generateSize,
	}
}
