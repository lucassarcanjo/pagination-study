package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/lucassarcanjo/pagination-study/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB  *gorm.DB
	Gin *gin.Engine
}

func (s *Server) InitDb(dsn string) *Server {
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	s.DB = db

	return s
}

func (s *Server) InitGin() *Server {
	g := gin.Default()
	s.Gin = g

	return s
}

func (s *Server) Ready() bool {
	return s.DB != nil && s.Gin != nil
}

func (s *Server) SeedData(amount int) *Server {
	s.DB.AutoMigrate(&model.User{})

	users := []model.User{}

	for i := 0; i < amount; i++ {
		var u model.User
		gofakeit.Struct(&u)
		u.CreatedAt = time.Now()
		users = append(users, u)
	}

	result := s.DB.Create(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	s.DB.Save(&users)

	fmt.Printf("Seed data created, added %d new users\n\n", len(users))

	return s
}

func (s *Server) Start(ep string) error {
	if !s.Ready() {
		return errors.New("Server isn't ready: make sure to init db and gin")
	}

	if err := http.ListenAndServe(ep, s.Gin.Handler()); err != nil {
		log.Fatal(err)
	}

	return nil
}
