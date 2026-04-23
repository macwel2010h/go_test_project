package runServer

import (
	"fmt"
	"log"
	"net"
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
	logger.Logger.Info("Starting Server...")
	listner, err := net.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Printf("Can not connect to the server.Make sure the port:%s is not in use.", s.Addr)
	}
	log.Printf("The server stareted. Please visit http://localhost%s", s.Addr)

	err = http.Serve(listner, handler)
	if err != nil {

		fmt.Printf("An error occured while starting the Server : %v\n", err)
		fmt.Println("Please restart the server.")
		return
	}

}

func FileServer() http.Handler {
	fileServer := http.FileServer(http.Dir("./web/"))
	logger.Logger.Info("The file server started.")
	return fileServer
}
