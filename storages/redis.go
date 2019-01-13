package storages

import (
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	redisclient *redis.Client
}

func (s *Redis) Init() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, error := client.Ping().Result()
	log.Print(pong, error)
	// Output: PONG <nil>
	s.redisclient = client
	return error
}

// generates the short url code
func (s *Redis) Code() string {

	return "42"
}

func (s *Redis) Save(url string) string {
	code := s.Code()

	err := s.redisclient.Set(code, url, 0).Err()
	if err != nil {
		panic(err)
	}

	return code
}

func (s *Redis) Load(code string) (string, error) {
	url, err := s.redisclient.Get(code).Result()
	if err != nil {
		panic(err)
	}

	return string(url), err
}
