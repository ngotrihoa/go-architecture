package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IServer interface {
	Run() error
	MapRouters() error
}

type Server struct {
	db  *gorm.DB
	gin *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	return &Server{gin: gin.Default(), db: db}
}

func (s *Server) Run() error {
	s.MapRouters()
	if err := s.gin.Run(":3000"); err != nil {
		log.Fatalln(err)
	}

	return nil
}
