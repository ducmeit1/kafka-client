package kits

import (
	"context"
	"flag"
	"fmt"
	"github.com/ducmeit1/kafka-client/common/middlewares"
	"github.com/ducmeit1/kafka-client/common/transports"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	HTTPPort   = flag.Int("http-port", 8888, "http port")
	ConfigPath = flag.String("config-path", "config.toml", "config path")
	ConfigType = flag.String("config-type", "toml", "config type")
)

type Server struct {
	Name        string
	Port        int
	RoutePrefix string
	Routes      []transports.Route
	Middlewares []http.Handler
	OnClose     func()
}

func init() {
	if err := Init(); err != nil {
		panic(err)
	}
}

func Init() error {
	flag.Parse()
	InitLogger()
	err := LoadConfig()
	if err != nil {
		return err
	}
	return nil
}

func NewServer(name string, onClose func()) (*Server, error) {
	return &Server{
		Name:    name,
		Port:    *HTTPPort,
		OnClose: onClose,
	}, nil
}

func (s *Server) AddRoutes(transport transports.Transport) {
	s.RoutePrefix = transport.PathPrefix
	s.Routes = transport.Routes
}

func (s *Server) Run() error {
	if len(s.Routes) == 0 {
		return fmt.Errorf("routes was not set")
	}

	//Adding middlewares
	m := middlewares.NewMiddlewares()
	m.AddGlobalMiddlewares()
	m.AddMiddlewares(s.Middlewares...)
	m.AddRoutes(s.RoutePrefix, s.Routes)

	h := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: m.GetHandlers(),
	}

	//Graceful shutdown handle
	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)

		<-c

		//Run on close
		s.Close()
		//A interrupt signal has sent to us, let's shutdown server with gracefully
		log.Info("Stopping server...")

		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		if err := h.Shutdown(ctx); err != nil {
			log.Errorf("Graceful shutdown has failed with error: %s", err)
		}
		close(idleConnsClosed)
	}()

	go func() {
		log.Infof("Starting: %v listen server on port %d", s.Name, s.Port)
		if err := h.ListenAndServe(); err != http.ErrServerClosed {
			log.Errorf("Run server on port %d has failed with error: %s", s.Port, err)
		} else {
			log.Errorf("Server was closed by shutdown gracefully")
		}
	}()

	<-idleConnsClosed

	return nil
}

func (s *Server) Close() {
	//Run OnClose callback
	s.OnClose()
}
