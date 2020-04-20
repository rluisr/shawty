package storages

import (
	"github.com/rluisr/shawty/models"
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	redisClient *redis.Client
}

func (s *Redis) Init() error {
	config := models.NewConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	pong, error := client.Ping().Result()
	log.Print(pong, error)
	// Output: PONG <nil>
	s.redisClient = client
	return error
}

// generates the short url code
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
	url, err := s.redisClient.Get(code).Result()
	if err != nil {
		panic(err)
	}

	return string(url), err
}
