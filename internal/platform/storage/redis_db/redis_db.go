package redis_db

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis_db database")
	Ctx    = context.TODO()
)

func NewDatabase(address string, password string, db int) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	return &Database{
		Client: client,
	}, nil
}

func Get(rdb *Database, key string) (string, error) {
	res, err := rdb.Client.Get(Ctx, key).Result()
	if err != nil {
		return "", errors.New("key does not exist")
	}

	return res, err
}

func Set(rdb *Database, key string, value string) (bool, error) {
	err := rdb.Client.Set(Ctx, key, value, 0).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func RPush(rdb *Database, key string, value string) (bool, error) {
	err := rdb.Client.RPush(Ctx, key, value, 0).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}
