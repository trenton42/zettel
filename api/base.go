package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"Status": "OK", "Version": "1.1"})
}
