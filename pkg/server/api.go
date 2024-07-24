package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trenton42/zettel/pkg/types"
)

func (s *Server) GetZettel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	z, err := s.db.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "zettel not found", "id": strconv.Itoa(id)})
	}
	return c.JSON(http.StatusOK, z)
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

func (s *Server) UpdateZettel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var z types.Zettel
	var req struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	z.Title = req.Title
	z.Body = req.Body
	err = s.db.Update(id, z)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, map[string]string{"msg": "Updated"})
}

func (s *Server) DeleteZettel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = s.db.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found or something"})
	}
	return c.NoContent(http.StatusNoContent)
}
