package config

import (
	"log"

	"github.com/asrofilfachrulr/transaction-api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	db := ConnectDataBase()
	s.DB = db

	s.Router = gin.Default()

	routes.Attach(s.Router, "db", db)

	routes.InitAPI(s.Router)
	log.Println("Server initiated...")
}

func (s *Server) Run() {
	sqlDB, _ := s.DB.DB()
	defer sqlDB.Close()

	log.Println("Running server at 8080...")
	s.Router.Run()
}
