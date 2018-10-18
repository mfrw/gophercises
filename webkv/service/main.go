package service

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type Service struct {
	Name        string
	TTL         time.Duration
	RedisClient redis.UniversalClient
}

func New(addrs []string, ttl time.Duration) (*Service, error) {
	s := new(Service)
	s.Name = "webkv"
	s.TTL = ttl
	s.RedisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: addrs,
	})
	ok, err := s.Check()
	if !ok {
		return nil, err
	}
	return s, nil
}

func (s *Service) Check() (bool, error) {
	_, err := s.RedisClient.Ping().Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := 200
	key := strings.Trim(r.URL.Path, "/")
	val, err := s.RedisClient.Get(key).Result()
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		status = 404
	}
	fmt.Fprintf(w, val)
	log.Printf("url=\"%s\" remote=\"%s\" key=\"%s\" status=%d\n", r.URL, r.RemoteAddr, key, status)
}
