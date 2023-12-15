package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trenton42/zettel/models"
)

func (s *Server) RegisterZettel() {
	s.e.GET("/zettels", createList[models.Zettel](s.db))
	s.e.POST("/zettels", createCreate[models.Zettel](s.db))
}

func createList[M any, K interface {
	models.Item
	*M
}](db DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		out := make([]models.Item, 0)
		a := K(new(M))
		var err error
		i := db.All(a.GetType())
		for err == nil {
			a := K(new(M))
			err = i.Next(a)
			if err == nil {
				out = append(out, a)
			}
		}
		return c.JSON(http.StatusOK, models.NewListResponse(out))
	}
}

func createCreate[M any, K interface {
	models.Item
	*M
}](db DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := K(new(M))
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, models.NewError("badrequest", "You've done something wrong").SetDetail(err.Error()).ToResponse())
		}
		err := db.CreateItem(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.NewError("badrequest", "You've done something wrong").SetDetail(err.Error()).ToResponse())
		}
		return c.JSON(http.StatusOK, models.NewResponse(req))
	}
}
