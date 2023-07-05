package internal

import (
	"IpLimiter/config"
	"IpLimiter/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.AppConfig
	Engine *gin.Engine
}

func newEngine(config *config.AppConfig, rateLimiter *RateLimiterMiddleware) *gin.Engine {
	gin.SetMode(config.GinMode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(rateLimiter.handle())
	return engine
}

func InitializeServer() (*Server, error) {
	appConfig, err := config.NewAppConfig()
	if err != nil {
		return nil, err
	}

	limiter, err := pkg.NewLimiter("")
	if err != nil {
		return nil, err
	}

	rateLimiter := NewRateLimiterMiddleware(limiter)
	engine := newEngine(appConfig, rateLimiter)

	return &Server{
		appConfig,
		engine,
	}, nil
}

func (s *Server) registerRoutes() {
	s.Engine.GET("/", sayHi)
}

func (s *Server) Run() {
	s.registerRoutes()

	Addr := ":" + s.Config.Port
	fmt.Printf("%s listening... \n", Addr)
	s.Engine.Run(Addr)
}
