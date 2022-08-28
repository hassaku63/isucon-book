package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	users := NewUser(
		WithSequentialIdUsers(5),
		WithRandomIdUsers(5),
		WithSequentialIdUsers(5),
	)

	fmt.Printf("users count: %d\n", users.Len())

	users.ForEach(func(i int, u *User) {
		fmt.Printf("%02d: user id: %d\n", i+1, u.ID)
	})
}

type User struct {
	ID int
}

type Users struct {
	mu   sync.RWMutex
	list []*User
	dict map[int]*User
}

type UserOption func(users *Users)

func NewUser(opts ...UserOption) *Users {
	users := &Users{
		mu:   sync.RWMutex{},
		list: make([]*User, 0),
		dict: make(map[int]*User, 0),
	}

	for _, opt := range opts {
		opt(users)
	}

	return users
}

func (u *Users) Add(user *User) (ok bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	if user.ID <= 0 {
		return false
	}

	// ID が重複するなら何もしない
	if _, found := u.dict[user.ID]; found {
		return false
	}

	u.list = append(u.list, user)
	u.dict[user.ID] = user
	return true
}

func (u *Users) Len() int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return len(u.list)
}

func (u *Users) ForEach(f func(i int, u *User)) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for i, user := range u.list {
		f(i, user)
	}
}

func (u *Users) MaxID() int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	maxID := 0
	for _, user := range u.list {
		if user.ID > maxID {
			maxID = user.ID
		}
	}

	return maxID
}

func WithSequentialIdUsers(count int) UserOption {
	return func(u *Users) {
		for i := 0; i < count; i++ {
			id := u.MaxID() + 1
			user := &User{
				ID: id,
			}

			u.Add(user)
		}
	}
}

func WithRandomIdUsers(count int) UserOption {
	return func(u *Users) {
		for i := 0; i < count; i++ {
			id := u.MaxID() + rand.Intn(101) - 50
			user := &User{
				ID: id,
			}

			ok := u.Add(user)
			if !ok {
				i--
			}
		}
	}
}
