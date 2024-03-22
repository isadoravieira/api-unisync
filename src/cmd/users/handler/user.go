package handler

import "net/http"

func StoreUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func IndexUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List user"))
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Query user by ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}