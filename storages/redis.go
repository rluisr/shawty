package storages

import (
	//"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

type Redis struct {
	redisclient *redis.Client
}

func (s *Redis) Init(c DbCredentials) error {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + strconv.Itoa(c.Port),
		Password: c.Pass,
		DB:       c.Name,
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
	// wenn key not found err: "redis: nil"
	if err != nil {
		//fmt.Println("%v", err)
		//s := err.Error()
		//log.Print("Fehler lesen code: " + code + " aus datenbank: " + s)
		log.Print("Fehler lesen code: " + code + " aus datenbank: " + err.Error())
		//url = ""
		//panic(err)
	}

	return string(url), err
}
