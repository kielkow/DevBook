package models

import (
	"time"
)

// User struct
type User struct {
	ID           uint64         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Nick         string         `json:"nick"`
	CreatedAt    time.Time      `json:"createdAt"`
	Followers    []User         `json:"followers"`
	Following    []User         `json:"following"`
	Publications []Publication `json:"publications"`
}
