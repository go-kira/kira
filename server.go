package kira

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

// StartServer - Start kira server
func (a *App) StartServer() {
	// define the server
	server := &http.Server{
		Addr:    a.Configs.GetString("SERVER_HOST") + ":" + strconv.Itoa(a.Configs.GetInt("SERVER_PORT")),
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
func (a *App) StartTLSServer() {
	// TODO:
	server := &http.Server{
		Addr:    a.Configs.GetString("SERVER_HOST") + ":" + strconv.Itoa(a.Configs.GetInt("SERVER_PORT")),
		Handler: a.Router,
	}

	// Gracefully shutdown
	go a.GracefullyShutdown(server)

	// Start server
	a.Log.Infof("Starting HTTP server, Listening at %q \n", "https://"+server.Addr)
	if err := server.ListenAndServeTLS("storage/framework/cert/server.crt", "storage/framework/cert/server.key"); err != http.ErrServerClosed {
		a.Log.Errorf("%v", err)
	} else {
		a.Log.Infof("Server closed!")
	}
}

// GracefullyShutdown the server
func (a *App) GracefullyShutdown(server *http.Server) {
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, os.Kill)

	sig := <-sigquit
	a.Log.Infof("Caught sig: %+v", sig)
	a.Log.Infof("Gracefully shutting down server...")

	if err := server.Shutdown(context.Background()); err != nil {
		a.Log.Fatalf("Unable to shutdown server: %v", err)
	} else {
		a.Log.Infof("Server stopped")
	}
}

// IsTLS - set the tls to true.
func (a *App) IsTLS() {
	a.isTLS = true
}
