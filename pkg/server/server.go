package server

import (
	"github.com/labstack/echo/v4"
	"github.com/trenton42/zettel/pkg/types"
)

type Server struct {
	*echo.Echo
	db types.DB
}

func New(db types.DB) *Server {
	s := &Server{
		Echo: echo.New(),
		db:   db,
	}
	s.registerRoutes()
	return s
}
