package main

import (
	"net/http"
	handlers "serv-test/api"
	"serv-test/middlewares"
	runServer "serv-test/server"
)

func RoutHandlers() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /web/", http.StripPrefix("/web", runServer.FileServer()))

	mux.HandleFunc("GET /{$}", handlers.IndexHandler)
	mux.HandleFunc("GET /signIn", handlers.SignInHandler)
	mux.HandleFunc("GET /about", handlers.AboutHandler)
	mux.HandleFunc("GET /create-account", handlers.CreateAccountHandler)

	mux.HandleFunc("POST /create-account", handlers.CreateUser)
	mux.HandleFunc("POST /signIn", handlers.PostSignInHandler)
	mux.HandleFunc("POST /create-post", handlers.PostHandler)

	return middlewares.LogRequest(middlewares.CommonHeaders(mux))
}
