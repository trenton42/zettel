package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trenton42/zettel/pkg/types"
)

func (s *Server) GetZettel(c echo.Context) error {
	return nil
}

func (s *Server) ListZettel(c echo.Context) error {
	z, err := s.db.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, z)
}

func (s *Server) CreateZettel(c echo.Context) error {
	var z types.Zettel
	var req struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	z.Title = req.Title
	z.Body = req.Body
	err = s.db.Create(z)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, map[string]string{"msg": "Created"})
}
