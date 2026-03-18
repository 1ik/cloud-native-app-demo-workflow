package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	hitCounterKey = "hitcount:counter"
	lockKey       = "hitcount:lock"
	lockTTL       = 5 * time.Second
	lockRetries   = 5
	lockBackoff   = 100 * time.Millisecond
)

// newRedisClient creates a Redis client using REDIS_HOST env (default redis:6379).
func newRedisClient() *redis.Client {
	host := getEnv("REDIS_HOST", "redis:6379")
	return redis.NewClient(&redis.Options{
		Addr: host,
	})
}

// getEnv returns env var if set, otherwise defaultVal.
func getEnv(key, defaultVal string) string {
	if v := lookupEnv(key); v != "" {
		return v
	}
	return defaultVal
}

// lookupEnv allows overriding in tests.
var lookupEnv = func(key string) string {
	if v, ok := sysLookupEnv(key); ok {
		return v
	}
	return ""
}

var sysLookupEnv = os.LookupEnv
var osHostname = os.Hostname

// acquireLock tries to obtain a lock with retries.
func acquireLock(ctx context.Context, rdb *redis.Client, key, value string, ttl time.Duration) error {
	for i := 0; i < lockRetries; i++ {
		ok, err := rdb.SetNX(ctx, key, value, ttl).Result()
		if err != nil {
			return err
		}
		if ok {
			log.Printf("lock acquired key=%s val=%s", key, value)
			return nil
		}
		time.Sleep(lockBackoff)
	}
	return errors.New("lock not acquired")
}

// releaseLock releases the lock only if owned by this holder.
func releaseLock(ctx context.Context, rdb *redis.Client, key, value string) {
	script := redis.NewScript(`
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end
    `)
	_ = script.Run(ctx, rdb, []string{key}, value).Err()
}

// HitHandler increments a shared counter with distributed locking and returns hostname + count.
func HitHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// unique lock value
		buf := make([]byte, 16)
		_, _ = rand.Read(buf)
		lockVal := hex.EncodeToString(buf)

		if err := acquireLock(ctx, rdb, lockKey, lockVal, lockTTL); err != nil {
			c.JSON(503, gin.H{"error": "busy, try again"})
			return
		}
		defer releaseLock(ctx, rdb, lockKey, lockVal)

		count, err := rdb.Incr(ctx, hitCounterKey).Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "redis error"})
			return
		}

		hostname, _ := osHostname()
		c.JSON(200, gin.H{
			"count":    count,
			"hostname": hostname,
		})
	}
}
