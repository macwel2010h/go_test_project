package main

import (
	"net/http"
	handlers "serv-test/api"
	"serv-test/config"
	"serv-test/middlewares"
	runServer "serv-test/server"
)

func RoutHandlers() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /web/", http.StripPrefix("/web", runServer.FileServer()))

	mux.Handle("GET /{$}", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.IndexHandler)))
	mux.Handle("GET /signIn", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.SignInHandler)))
	mux.Handle("GET /about", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.AboutHandler)))
	mux.Handle("GET /create-account", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.CreateAccountHandler)))

	mux.Handle("POST /create-account", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /signIn", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.PostSignInHandler)))
	mux.Handle("POST /create-post", config.App.SessionManager.LoadAndSave(http.HandlerFunc(handlers.PostHandler)))

	return middlewares.LogRequest(middlewares.CommonHeaders(mux))
}
