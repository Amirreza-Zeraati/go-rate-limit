package initializers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var Client *redis.Client
var Ctx = context.Background()

func ConnectToDB() {

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s:%s", host, port)
	Client = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})
}
