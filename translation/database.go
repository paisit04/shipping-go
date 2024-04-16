package translation

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/paisit04/shipping-go/config"
	"github.com/paisit04/shipping-go/handlers/rest"
)

var _ rest.Translator = &Database{}

type Database struct {
	conn *redis.Client
}

func NewDatabaseService(cfg config.Configuration) *Database {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DatabaseURL, cfg.DatabasePort),
		Password: "",
		DB:       0,
	})
	return &Database{
		conn: rdb,
	}
}

func (s *Database) Close() error {
	return s.conn.Close()
}

func (s *Database) Translate(word, language string) string {
	out := s.conn.Get(fmt.Sprintf("%s:%s", word, language))
	return out.Val()
}
