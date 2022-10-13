package delivery

import (
	"fmt"

	"enigmacamp.com/golang-sample/config"
	"enigmacamp.com/golang-sample/delivery/controller"
	"enigmacamp.com/golang-sample/manager"
	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}
func (s *Server) initController() {
	controller.NewCustomerController(s.engine, s.useCaseManager.CustomerUseCase())
}
func NewServer() *Server {
	c := config.NewConfig()
	r := gin.Default()
	infra := manager.NewInfraManager(c)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	if c.ApiHost == "" || c.ApiPort == "" {
		panic("No Host or port define")
	}
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{useCaseManager: usecase, engine: r, host: host}
}
