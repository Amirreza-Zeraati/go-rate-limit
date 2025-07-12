package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	ini "go-rate-limit/initializers"
	"net/http"
	"time"
)

const (
	limitPeriod = 30 * time.Second
	limitCount  = 5
)

func init() {
	ini.LoadEnvFile()
	ini.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", rateLimitMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	err := r.Run()
	if err != nil {
		return
	}
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "IP : " + ip

		count, err := ini.Client.Get(ini.Ctx, key).Int()
		if err != nil && !errors.Is(err, redis.Nil) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		if count >= limitCount {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "rate limit exceeded",
				"message": fmt.Sprintf("please wait %s", limitPeriod),
			})
			return
		}

		_, err = ini.Client.Incr(ini.Ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		if count == 0 {
			ini.Client.Expire(ini.Ctx, key, limitPeriod)
		}
		c.Next()
	}
}
