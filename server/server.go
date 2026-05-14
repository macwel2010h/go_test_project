package runServer

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"serv-test/config"
	logger "serv-test/log"
	"time"
)

func RunServer(handler http.Handler) {
	var s = http.Server{
		Addr:         ":8001",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
		ErrorLog:     slog.NewLogLogger(logger.Logger.Handler(), slog.LevelError),
		TLSConfig:    config.TlsConfig,
	}
	logger.Logger.Info("Starting Server...")
	listner, err := net.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Printf("Can not connect to the server.Make sure the port:%s is not in use.", s.Addr)
	}
	log.Printf("The server stareted. Please visit https://localhost%s", s.Addr)

	err = s.ServeTLS(listner, "tls/cert.pem", "tls/key.pem")
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
