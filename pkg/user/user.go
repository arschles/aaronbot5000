package user

import (
	"sync"

	lru "github.com/hashicorp/golang-lru"
)

func init() {
	users, _ = lru.New(100)
	twitchUsers, _ = lru.New(100)
}

var users *lru.Cache
var twitchUsers *lru.Cache

var userID string
var mainChannel string

type User struct {
	ID          string         `json:"id"`
	DisplayName string         `json:"displayName"`
	Color       string         `json:"color"`
	Badges      map[string]int `json:"badges"`
	Points      uint64         `json:"points"`
	New         bool           `json:"-"`
	IsFollower  bool           `json:"isFollower"`
	lock        sync.RWMutex
}

// func (u *User) GivePoints(points uint64) error {
// 	u.lock.Lock()
// 	defer u.lock.Unlock()

// 	u.Points = u.Points + points
// 	return updateUser(u)
// }

// func (u *User) TakePoints(points uint64) error {
// 	u.lock.Lock()
// 	defer u.lock.Unlock()

// 	u.Points = u.Points - points
// 	return updateUser(u)
// }

// func (u *User) TransferPoints(points uint64, userID string) error {
// 	u.lock.Lock()
// 	defer u.lock.Unlock()

// 	u2, err := GetUser(userID)
// 	if err != nil {
// 		return err
// 	}

// 	// If insufficient balance. Transfer remaining balance.
// 	if u.Points < points {
// 		points = u.Points
// 	}
// 	u.Points = u.Points - points

// 	u2.Points = u2.Points + points

// 	jsonUser1, err := json.Marshal(u)
// 	if err != nil {
// 		return err
// 	}

// 	jsonUser2, err := json.Marshal(u2)
// 	if err != nil {
// 		return err
// 	}
// 	return db.Update(func(tx *bbolt.Tx) error {
// 		bucket := tx.Bucket(USER_BUCKET)
// 		if err := bucket.Put([]byte(u.ID), jsonUser1); err != nil {
// 			return err
// 		}
// 		if err := bucket.Put([]byte(u2.ID), jsonUser2); err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }

// func (u *User) Save() error {
// 	u.lock.Lock()
// 	defer u.lock.Unlock()

// 	return updateUser(u)
// }

// func updateUser(u *User) error {
// 	u.New = false

// 	buf, err := json.Marshal(u)
// 	if err != nil {
// 		return err
// 	}

// 	return db.Update(func(tx *bbolt.Tx) error {
// 		bucket := tx.Bucket(USER_BUCKET)
// 		err := bucket.Put([]byte(u.ID), buf)
// 		return err
// 	})
// }

// func GetUser(id string) (*User, error) {
// 	var u User

// 	if u, ok := users.Get(id); ok {
// 		return u.(*User), nil
// 	}

// 	err := db.View(func(tx *bbolt.Tx) error {
// 		b := tx.Bucket(USER_BUCKET)
// 		v := b.Get([]byte(id))
// 		if len(v) == 0 {
// 			u.ID = id
// 			u.New = true
// 			return nil
// 		}

// 		return json.Unmarshal(v, &u)
// 	})

// 	if err == nil {
// 		users.Add(id, &u)
// 	}

// 	return &u, err
// }

// func GetUserByName(name string) (*User, error) {
// 	if u, ok := twitchUsers.Get(name); ok {
// 		return GetUser(u.(helix.User).ID)
// 	}

// 	resp, err := helixClient.GetUsers(&helix.UsersParams{
// 		Logins: []string{name},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(resp.Data.Users) == 0 {
// 		return nil, fmt.Errorf("User with name '%s' was not found.", name)
// 	}

// 	twitchUsers.Add(name, resp.Data.Users[0])
// 	return GetUser(resp.Data.Users[0].ID)
// }

// func getMainChannel() string {
// 	if mainChannel != "" {
// 		return mainChannel
// 	}

// 	type twitchConfig struct {
// 		MainChannel string `json:"mainChannel"`
// 	}

// 	if c, ok := config.ModuleConfig["twitch"]; ok {
// 		var tc twitchConfig
// 		if err := json.Unmarshal(c, &tc); err == nil {
// 			mainChannel = tc.MainChannel
// 		}
// 	}

// 	return mainChannel
// }

// func getUserID() string {
// 	if userID != "" {
// 		return userID
// 	}

// 	if u, err := GetUserByName(getMainChannel()); err == nil {
// 		userID = u.ID
// 	}

// 	return userID
// }
