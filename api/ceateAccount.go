package handlers

import "net/http"

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/createAccount.html")
}
