package kira

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// StartServer - Start kira server
func (a *App) StartServer() {
	// Server HOST/PORT
	host := a.Configs.GetString("server.host", "127.0.0.1")
	port := a.Configs.GetInt("server.port", 8080)

	// define the server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: a.Router,
	}

	// Gracefully shutdown
	go a.GracefullyShutdown(server)

	// Start server
	a.Log.Infof("Starting HTTP server, Listening at %q \n", "http://"+server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		a.Log.Errorf("%v", err)
	} else {
		a.Log.Infof("Server closed!")
	}
}

// StartTLSServer - start an TLS server provided by: Let's Encrypt.
// To generate keys:
//  - openssl genrsa -out server.key 2048
//  - openssl ecparam -genkey -name secp384r1 -out server.key
//  - openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
func (a *App) StartTLSServer() {
	server := &http.Server{
		Addr:    a.Configs.GetString("server.host", "127.0.0.1") + ":" + strconv.Itoa(a.Configs.GetInt("server.port", 8080)),
		Handler: a.Router,
	}

	// Gracefully shutdown
	go a.GracefullyShutdown(server)

	// Start server
	a.Log.Infof("Starting HTTPS server, Listening at %q \n", "https://"+server.Addr)

	// Certificate & Key
	certificateFile := a.Configs.GetString("server.tls_certificate", "./server.crt")
	keyFile := a.Configs.GetString("server.tls_key", "./server.key")

	if err := server.ListenAndServeTLS(certificateFile, keyFile); err != http.ErrServerClosed {
		a.Log.Errorf("%v", err)
	} else {
		a.Log.Infof("Server closed!")
	}
}

// GracefullyShutdown the server
func (a *App) GracefullyShutdown(server *http.Server) {
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)

	sig := <-sigquit
	a.Log.Infof("Signal to shutdown the server: %+v", sig)

	if err := server.Shutdown(context.Background()); err != nil {
		a.Log.Fatalf("Unable to shutdown server: %v", err)
	}
}
