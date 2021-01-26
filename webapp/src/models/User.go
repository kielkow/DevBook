package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
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

	return User{}, nil
}

// SearchUserData func
func SearchUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if error = json.NewDecoder(response.Body).Decode(&user); error != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// SearchFollowers func
func SearchFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if error = json.NewDecoder(response.Body).Decode(&followers); error != nil {
		channel <- nil
		return
	}

	channel <- followers
}

// SearchFollowing func
func SearchFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if error = json.NewDecoder(response.Body).Decode(&following); error != nil {
		channel <- nil
		return
	}

	channel <- following
}

// SearchPublications func
func SearchPublications(channel chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publication
	if error = json.NewDecoder(response.Body).Decode(&publications); error != nil {
		channel <- nil
		return
	}

	channel <- publications
}
