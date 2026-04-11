package main

import (
	"net/http"
	handlers "serv-test/api"
)

func RoutHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/signIn", handlers.SignInHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/create-account", handlers.CreateAccountHandler)

	return mux
}
