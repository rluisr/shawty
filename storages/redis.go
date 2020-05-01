package storages

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/rluisr/shawty/lib"
	"github.com/rluisr/shawty/models"
	"strconv"
	"time"
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
			DB:            config.RedisDB,
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

func (s *Redis) Save(url, expire string) (string, error) {
	var code string

	for {
		code = s.Code()

		err := s.redisClient.Get(code).Err()
		if errors.Is(err, redis.Nil) {
			expireInt, err := strconv.Atoi(expire)
			if err != nil {
				return "", err
			}

			err = s.redisClient.Set(code, url, time.Duration(expireInt)*24*time.Hour).Err()
			if err != nil {
				return "", err
			}
			break
		}
	}

	return code, nil
}

func (s *Redis) Load(code string) (string, error) {
	return s.redisClient.Get(code).Result()
}
