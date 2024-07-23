package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerRoutes() {
	g := s.Group("/api/v1")
	g.GET("/zettel/:id", s.GetZettel)
	g.GET("/zettel/", s.ListZettel)
	g.POST("/zettel/", s.CreateZettel)

	s.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It ok.")
	})
}
