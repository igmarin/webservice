package controllers

import "net/http"

// RegisterControllers my main controller for userController
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
