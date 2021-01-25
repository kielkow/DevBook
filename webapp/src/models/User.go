package models

import (
	"net/http"
	"time"
)

// User struct
type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreatedAt    time.Time     `json:"createdAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

// SearchCompletedUser func
func SearchCompletedUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go SearchUserData(userChannel, userID, r)
	go SearchFollowers(followersChannel, userID, r)
	go SearchFollowing(followingChannel, userID, r)
	go SearchPublications(publicationsChannel, userID, r)
}

// SearchUserData func
func SearchUserData(channel <-chan User, userID uint64, r *http.Request) {

}

// SearchFollowers func
func SearchFollowers(channel <-chan []User, userID uint64, r *http.Request) {

}

// SearchFollowing func
func SearchFollowing(channel <-chan []User, userID uint64, r *http.Request) {

}

// SearchPublications func
func SearchPublications(channel <-chan []Publication, userID uint64, r *http.Request) {

}
