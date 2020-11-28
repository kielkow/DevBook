package controllers

import "net/http"

// CreateUser func
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a user"))
}

// SearchUsers func
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching users"))
}

// SearchUser func
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching a user"))
}

// UpdateUser func
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser func
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user"))
}
