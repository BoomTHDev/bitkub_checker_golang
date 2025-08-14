package server

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/boomthdev/wld_check_bk/config"
	"github.com/boomthdev/wld_check_bk/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type fiberServer struct {
	app  *fiber.App
	conf *config.Config
}

var (
	once           sync.Once
	serverInstance *fiberServer
)

func NewFiberServer(conf *config.Config) *fiberServer {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit:    conf.Server.BodyLimit,
		IdleTimeout:  time.Second * time.Duration(conf.Server.TimeOut),
		ErrorHandler: middleware.ErrorHandler(),
	})

	once.Do(func() {
		serverInstance = &fiberServer{
			app:  fiberApp,
			conf: conf,
		}
	})

	return serverInstance
}

func (s *fiberServer) setupRoutes() {
	s.app.Use(logger.New())
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(s.conf.Server.AllowOrigins, ","),
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-BITKUB-API-KEY, X-BITKUB-API-SECRET",
	}))

	s.initWalletRouter()

	s.app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Sorry, endpoint %s %s not found", ctx.Method(), ctx.Path()),
		})
	})
}

func (s *fiberServer) Start() {
	s.setupRoutes()
	s.httpListening()
}

func (s *fiberServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Listen(url); err != nil {
		fmt.Printf("Error: %s", err)
	}
}

func BuildVercelHandler(conf *config.Config) http.Handler {
	server := NewFiberServer(conf)
	server.setupRoutes()
	return adaptor.FiberApp(server.app)
}
