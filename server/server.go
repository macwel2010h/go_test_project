package runServer

import (
	"fmt"
	"log"
	"net/http"
	logger "serv-test/log"
	"time"
)

var s = http.Server{
	Addr:         ":8001",
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
	IdleTimeout:  120 * time.Second,
}

func RunServer(handler http.Handler) {

	err := http.ListenAndServe(s.Addr, handler)
	if err != nil {

		fmt.Printf("An error occured while starting the Server : %v\n", err)
		fmt.Println("Please restart the server.")
		return
	}
	logger.Logger.Info("Starting Server...")
	log.Printf("The server stareted. Please visit http://localhost%s", s.Addr)
}

func FileServer() http.Handler {
	fileServer := http.FileServer(http.Dir("./web/"))
	logger.Logger.Info("The file server started.")
	return fileServer
}
