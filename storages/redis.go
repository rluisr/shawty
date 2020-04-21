package storages

import (
	"github.com/go-redis/redis"
	"github.com/rluisr/shawty/lib"
	"github.com/rluisr/shawty/models"
)

var config = models.NewConfig()

type Redis struct {
	redisClient *redis.Client
}

func (s *Redis) Init() error {
	config := models.NewConfig()

	if len(config.RedisSentinelAddr) == 0 {
		s.redisClient = redis.NewClient(&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPassword,
			DB:       config.RedisDB,
		})
	} else {
		s.redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    config.RedisSentinelMasterName,
			Password:      config.RedisPassword,
			SentinelAddrs: config.RedisSentinelAddr,
		})
	}

	_, err := s.redisClient.Ping().Result()
	return err
}

// generate
func (s *Redis) Code() string {
	return lib.RandString(config.GenerateSize)
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
