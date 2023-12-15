package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/trenton42/zettel/models"
)

func (s *Server) login(c echo.Context) error {
	r := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusUnauthorized, models.NewError("notauthorized", "Not Authorized").SetDetail(err.Error()).ToResponse())
	}

	t, err := s.db.GetToken(r.Username, r.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.NewError("notauthorized", "Not Authorized").SetDetail(err.Error()).ToResponse())
	}
	return c.JSON(http.StatusOK, models.NewResponse(&t))
}

func (s *Server) createUser(c echo.Context) error {
	r := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("BadRequest", "Username and Password are required"))
	}

	u, err := s.db.CreateUser(r.Username, r.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("BadRequest", "You did something wrong"))
	}

	return c.JSON(http.StatusOK, models.NewResponse(&u))
}

func (s *Server) getUser(c echo.Context) error {
	email := c.Param("email")
	var u models.User
	err := s.db.GetItem(email, &u)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.NewErrorResponse("NotFound", fmt.Sprintf("User %s was not found", email)))
	}
	return c.JSON(http.StatusOK, models.NewResponse(&u))
}

func (s *Server) getUsers(c echo.Context) error {
	out := make([]models.Item, 0)
	i := s.db.All((&models.User{}).GetType())
	var err error
	for err == nil {
		m := models.User{}
		err = i.Next(&m)
		if err == nil {
			out = append(out, &m)
		} else {
			fmt.Printf("========> error: %s\n", err)
		}
	}
	return c.JSON(http.StatusOK, models.NewListResponse(out))
}

func (s *Server) requireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, models.NewErrorResponse("unauthorized", "You must have a valid token to access this resource"))
		}
		ts := strings.SplitN(token, " ", 2)
		if len(ts) != 2 || ts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, models.NewErrorResponse("unauthorized", "You must have a valid token to access this resource"))
		}
		token = ts[1]
		tobj := models.Token{}
		err := s.db.QueryOne("key", "==", token, &tobj)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.NewError("unauthorized", "You do not have access").SetDetail(err.Error()).ToResponse())
		}
		c.Response().Header().Set(echo.HeaderServer, token)
		return next(c)
	}
}
