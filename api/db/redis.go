package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/lucas-code42/url-shortner/configs"
)

func MountRedis() Cache {
	// consult docker-compose
	var rds Cache = NewCacheDb(redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.REDIS_HOST, configs.REDIS_PORT),
		Password: configs.REDIS_PASSWORD,
		DB:       0, // default db
	}))
	return rds
}

type Cache interface {
	Get(alias string) (string, error)
	Create(url, urlAlias string) error
	CloseDb()
}

type CacheDb struct {
	Db *redis.Client
}

func NewCacheDb(redisClient *redis.Client) *CacheDb {
	return &CacheDb{Db: redisClient}
}

func (c *CacheDb) Create(url, urlAlias string) error {
	result := c.Db.Set(urlAlias, url, time.Duration(time.Minute*10))
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (c *CacheDb) Get(alias string) (string, error) {
	result := c.Db.Get(alias)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Val(), nil
}

func (c *CacheDb) CloseDb() {
	c.Db.Close()
}
