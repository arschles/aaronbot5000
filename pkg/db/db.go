package db

// punchypenguin: Russians robots same thing
// YOU DON'T PLAY CHESS, CHESS PLAYS YOU

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nicklaw5/helix"
)

// var db *bbolt.DB
var USER_BUCKET = []byte("Users")
var FOLLOWER_BUCKET = []byte("Followers")
var COUNTER_BUCKET = []byte("Counters")

// func IncrementCounter(counter string) (current uint64) {
// 	db.Update(func(tx *bbolt.Tx) error {
// 		b := tx.Bucket(COUNTER_BUCKET)
// 		v := b.Get([]byte(counter))

// 		if len(v) >= 1 {
// 			if err := json.Unmarshal(v, &current); err != nil {
// 				return err
// 			}
// 		}
// 		current++

// 		j, err := json.Marshal(current)
// 		if err != nil {
// 			return err
// 		}
// 		b.Put([]byte(counter), j)

// 		return nil
// 	})

// 	return
// }

// func ListCounters() []string {
// 	counters := make([]string, 0)

// 	db.View(func(tx *bbolt.Tx) error {
// 		b := tx.Bucket(COUNTER_BUCKET)

// 		b.ForEach(func(k, v []byte) error {
// 			counters = append(counters, string(k))
// 			return nil
// 		})
// 		return nil
// 	})

// 	return counters
// }

// New creates a new redis client with the given address and password
func New(ctx context.Context) (*redis.Client, error) {
	redisAddr := os.Getenv("BOT_REDIS_ADDR")
	redisPass := os.Getenv("BOT_REDIS_PASS")
	if redisAddr == "" {
		return nil, errors.New("BOT_REDIS_ADDR missing")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

// StartUpdateFollowersLoop starts a loop in a goroutine to periodically
// update followers from the twitch API
func StartUpdateFollowersLoop(
	ctx context.Context,
	followTarget string,
	redisCl *redis.Client,
	helixCl *helix.Client,
) {
	go func() {
		UpdateFollowers(ctx, followTarget, redisCl, helixCl)
		t := time.NewTicker(5 * time.Minute)
		for range t.C {
			UpdateFollowers(ctx, followTarget, redisCl, helixCl)
		}
	}()
}
