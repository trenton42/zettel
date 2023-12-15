package data

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/trenton42/zettel/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) CreateUser(username, password string) (u models.User, err error) {
	if username == "" || password == "" {
		err = fmt.Errorf("username and password must not be blank")
		return
	}
	u.Email = username
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.Password = string(p)
	err = s.CreateIndexedItem(&u)
	return
}

func (s *Store) GetToken(username, password string) (t models.Token, err error) {
	var u models.User
	err = s.GetIndexedItem(username, &u)

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return
	}
	tRef := s.client.Collection("tokens").NewDoc()
	b := make([]byte, 15)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	w := &strings.Builder{}
	wc := base64.NewEncoder(base64.StdEncoding, w)
	wc.Write(b)
	wc.Close()
	t.SetID(tRef.ID)
	t.Key = w.String()
	t.User = username
	t.TTL = time.Now().Add(time.Hour * 48)
	_, err = tRef.Set(context.Background(), &t)
	return
}
