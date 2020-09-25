package db

// punchypenguin: Russians robots same thing
// YOU DON'T PLAY CHESS, CHESS PLAYS YOU

import (
	"time"

	"github.com/go-redis/redis/v8"
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

func New(redisAddr, redisPass string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	go func() {
		UpdateFollowers(rdb)
		t := time.NewTicker(5 * time.Minute)
		for range t.C {
			UpdateFollowers(rdb)
		}
	}()

	return rdb, nil
}
