package data

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/trenton42/zettel/models"
)

type Store struct {
	client *firestore.Client
}

func New(projectID string) (*Store, error) {
	s := &Store{}
	var err error
	s.client, err = firestore.NewClient(context.Background(), projectID)
	return s, err
}

func (s *Store) Close() {
	if s.client != nil {
		s.client.Close()
	}
}

type Iterator struct {
	i *firestore.DocumentIterator
}

func (i *Iterator) Next(item models.Item) error {
	c, err := i.i.Next()
	if err != nil {
		return err
	}
	item.SetID(c.Ref.ID)
	return c.DataTo(item)
}
