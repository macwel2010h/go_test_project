package handlers

import "net/http"

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/signIn.html")
}
