package main

import (
	"net/http"
	handlers "serv-test/api"
	"serv-test/config"
	"serv-test/middlewares"
	runServer "serv-test/server"

	"github.com/justinas/alice"
)

func RoutHandlers() http.Handler {

	mux := http.NewServeMux()

	SessionMan := alice.New(config.App.SessionManager.LoadAndSave)

	mux.Handle("GET /web/", http.StripPrefix("/web", runServer.FileServer()))

	mux.Handle("GET /{$}", SessionMan.ThenFunc(handlers.IndexHandler))
	mux.Handle("GET /signIn", SessionMan.ThenFunc(handlers.SignInHandler))
	mux.Handle("GET /about", SessionMan.ThenFunc(handlers.AboutHandler))
	mux.Handle("GET /create-account", SessionMan.ThenFunc(handlers.CreateAccountHandler))
	mux.Handle("GET /welcome", SessionMan.ThenFunc(handlers.WelcomeHandler))

	mux.Handle("POST /create-account", SessionMan.ThenFunc(handlers.CreateUser))
	mux.Handle("POST /signIn", SessionMan.ThenFunc(handlers.PostSignInHandler))
	mux.Handle("POST /create-post", SessionMan.ThenFunc(handlers.PostHandler))

	PanicLogHeaders := alice.New(middlewares.PanicRecover, middlewares.LogRequest, middlewares.CommonHeaders)

	return PanicLogHeaders.Then(mux)

}
