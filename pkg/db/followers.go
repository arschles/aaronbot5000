package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arschles/aaronbot5000/pkg/user"
	"github.com/go-redis/redis/v8"
	"github.com/nicklaw5/helix"
)

func isFollower(ctx context.Context, redisCl *redis.Client, userID string) (bool, error) {
	follower, err := redisCl.HGet(ctx, "followers", userID).Result()
	if err != nil {
		return false, err
	}
	return follower != "", nil
}

// UpdateFollowers updates followers in the database
func UpdateFollowers(
	ctx context.Context,
	followTarget string,
	rdb *redis.Client,
	helixCl *helix.Client,
) {
	const batchSize = 100
	fmt.Println("Update of followers started.")
	defer fmt.Println("Update of followers finished.")

	// update the followers set
	cursor := ""
	for {
		resp, err := helixCl.GetUsersFollows(&helix.UsersFollowsParams{
			After: cursor,
			First: batchSize,
			ToID:  followTarget,
		})
		if err != nil {
			log.Printf("Error getting followers")
			continue
		}

		for _, f := range resp.Data.Follows {
			j, err := json.Marshal(f)
			if err != nil {
				log.Printf("Could not marshal follows data for user %s", f.FromName)
				continue
			}

			if err := rdb.HMSet(ctx, f.FromID, j).Err(); err != nil {
				log.Printf("Error adding follower '%s' to DB (%s)", f.FromName, err)
				continue
			}
		}

		// bail out if we are on the last page, since we're getting
		// batches of 100 each loop iteration
		if len(resp.Data.Follows) < batchSize {
			break
		}
		cursor = resp.Data.Pagination.Cursor
	}

	// update the users set to mark new followers as followers
	allUsers, err := rdb.HVals(ctx, "users").Result()
	if err != nil {
		log.Printf("Couldn't get list of users from DB (%s)", err)
		return
	}

	for _, userStr := range allUsers {
		var u user.User
		err := json.Unmarshal([]byte(userStr), &u)
		if err != nil {
			log.Printf("Couldn't unmarshal user (%s)", err)
			continue
		}
		// Check Followers bucket to see if this id exists
		follower, err := isFollower(ctx, rdb, u.ID)
		if err != nil {
			log.Printf("Couldn't figure out if user %s is a follower", u.ID)
			continue
		}
		u.IsFollower = follower
		if err := saveUser(ctx, rdb, &u); err != nil {
			log.Printf("Couldn't save user %v (%s)", u, err)
			continue
		}
	}
}

func saveUser(ctx context.Context, redisCl *redis.Client, user *user.User) error {
	var userBytes []byte
	var err error
	if userBytes, err = json.Marshal(&user); err != nil {
		return err
	}
	return redisCl.HSet(ctx, "users", map[string]interface{}{
		user.ID: userBytes,
	}).Err()
}
