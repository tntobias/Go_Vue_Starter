package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"
	eLog "github.com/neko-neko/echo-logrus"
	"github.com/neko-neko/echo-logrus/log"

	"github.com/tntobias/Go_Vue_Starter/config"
)

type Server struct {
	config *config.Config
	echo   *echo.Echo
}

func New(conf *config.Config) *Server {
	server := &Server{
		config: conf,
	}

	e := echo.New()
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.DEBUG)
	log.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	e.Logger = log.Logger()
	e.Use(eLog.Logger())
	e.GET("api/info", server.Info)
	e.Static("/", "client/dist")
	e.File("/", "client/dist/index.html")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	server.echo = e
	return server
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%s", s.config.Port))
}

func (s *Server) Info(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"description": "just getting started",
		"name":        "starter-project",
		"version":     "0.1.0",
	})
}
