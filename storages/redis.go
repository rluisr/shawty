package storages

import (
	"github.com/go-redis/redis"
	"github.com/rluisr/shawty/models"
)

type Redis struct {
	redisClient *redis.Client
}

func (s *Redis) Init() error {
	config := models.NewConfig()

	s.redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	_, err := s.redisClient.Ping().Result()
	return err
}

func (s *Redis) Code() string {
	return "42"
}

func (s *Redis) Save(url string) string {
	code := s.Code()

	err := s.redisClient.Set(code, url, 0).Err()
	if err != nil {
		panic(err)
	}

	return code
}

func (s *Redis) Load(code string) (string, error) {
	return s.redisClient.Get(code).Result()
}
