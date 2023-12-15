package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trenton42/zettel/models"
)

type Server struct {
	e  *echo.Echo
	db DB
}

type DB interface {
	GetToken(string, string) (models.Token, error)
	CreateUser(string, string) (models.User, error)
	GetItem(string, models.Item) error
	GetIndexedItem(string, models.Indexed) error
	CreateItem(models.Item) error
	QueryOne(string, string, string, models.Item) error
	Query(string, string, string, string) models.Iterator
	All(string) models.Iterator
}

func New(db DB) *Server {
	s := &Server{}
	s.db = db
	s.e = echo.New()
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.Logger())

	// Routes
	s.e.GET("/", s.root)
	s.e.POST("/auth/login", s.login)
	s.e.POST("/auth/user", s.createUser)
	s.e.GET("/auth/user", s.getUsers)
	s.e.GET("/auth/user/:email", s.getUser, s.requireLogin)
	s.RegisterZettel()
	return s
}

func (s *Server) Serve() error {
	return s.e.Start(":1123")
}
