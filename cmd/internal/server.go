package internal

import (
	"IpLimiter/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.AppConfig
	Engine *gin.Engine
}

func NewEngine(config *config.AppConfig) *gin.Engine {
	gin.SetMode(config.GinMode)

	engine := gin.Default()
	engine.Use(gin.Recovery())
	return engine
}

func NewServer(config *config.AppConfig, engine *gin.Engine) *Server {
	return &Server{
		config,
		engine,
	}
}

func InitializeServer() (*Server, error) {
	appConfig, err := config.NewAppConfig()
	if err != nil {
		return nil, err
	}

	engine := NewEngine(appConfig)
	server := NewServer(appConfig, engine)

	return server, nil
}

func (s *Server) registerRoutes() {
	s.Engine.GET("/", func(context *gin.Context) {
		fmt.Println("Hello")
	})
}

func (s *Server) Run() {
	s.registerRoutes()

	Addr := ":" + s.Config.Port
	fmt.Printf("%s listening... \n", Addr)
	s.Engine.Run(Addr)
}
