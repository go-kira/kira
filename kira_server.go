package kira

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kira/config"
)

// StartServer - Start kira server
func (app *App) StartServer(server *http.Server) {
	// Gracefully shutdown
	go app.GracefullyShutdown(server)

	// Start server
	app.logger.Infof("Starting HTTP server, Listening at %q \n", "http://"+server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		app.logger.Errorf("%v", err)
	} else {
		app.logger.Infof("Server closed!")
	}
}

// StartTLSServer - start an TLS server provided by: Let's Encrypt.
// To generate keys:
//  - openssl genrsa -out server.key 2048
//  - openssl ecparam -genkey -name secp384r1 -out server.key
//  - openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
func (app *App) StartTLSServer(server *http.Server) {
	// Gracefully shutdown
	go app.GracefullyShutdown(server)

	// Start server
	app.logger.Infof("Starting HTTPS server, Listening at %q \n", "https://"+server.Addr)

	// Certificate & Key
	certificateFile := app.Configs.GetString("server.tls_certificate", "./server.crt")
	keyFile := app.Configs.GetString("server.tls_key", "./server.key")

	if err := server.ListenAndServeTLS(certificateFile, keyFile); err != http.ErrServerClosed {
		app.logger.Errorf("%v", err)
	} else {
		app.logger.Infof("Server closed!")
	}
}

// GracefullyShutdown the server
func (app *App) GracefullyShutdown(server *http.Server) {
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)

	sig := <-sigquit
	app.logger.Infof("Signal to shutdown the server: %+v", sig)

	if err := server.Shutdown(context.Background()); err != nil {
		app.logger.Fatalf("Unable to shutdown server: %v", err)
	}
}

func serverAddr(config *config.Config, addr ...string) string {
	if len(addr) > 0 {
		return addr[0]
	}

	// Server HOST/PORT
	host := config.GetString("server.host", "127.0.0.1")
	port := config.GetInt("server.port", 8080)

	// Server Addr
	return fmt.Sprintf("%s:%d", host, port)
}
