package data

import (
	"context"
	"fmt"

	"github.com/trenton42/zettel/models"
)

func (db *Store) CreateItem(i models.Item) error {
	if val, ok := i.(models.Indexed); ok {
		return db.CreateIndexedItem(val)
	}
	ref := db.client.Collection(i.GetType()).NewDoc()
	_, err := ref.Create(context.Background(), i)
	if err != nil {
		return err
	}
	i.SetID(ref.ID)
	return nil
}

func (db *Store) CreateIndexedItem(i models.Indexed) error {
	if i.GetIndex() == "" {
		return fmt.Errorf("indexed field must not be empty")
	}
	ref := db.client.Collection(i.GetType()).NewDoc()
	idx := models.IndexRef{
		Type: i.GetType(),
		ID:   ref.ID,
	}
	_, err := db.client.Collection(fmt.Sprintf("index_%s", i.GetType())).Doc(i.GetIndex()).Create(context.Background(), &idx)
	if err != nil {
		return err
	}
	_, err = ref.Create(context.Background(), i)
	i.SetID(ref.ID)
	return err
}

func (db *Store) GetItem(id string, i models.Item) error {
	if idx, ok := i.(models.Indexed); ok {
		return db.GetIndexedItem(id, idx)
	}
	ref, err := db.client.Collection(i.GetType()).Doc(id).Get(context.Background())
	if err != nil {
		return err
	}
	i.SetID(ref.Ref.ID)
	return ref.DataTo(i)
}

func (db *Store) GetIndexedItem(id string, i models.Indexed) error {
	if id == "" {
		return fmt.Errorf("indexed field must not be empty")
	}
	ref, err := db.client.Collection(fmt.Sprintf("index_%s", i.GetType())).Doc(id).Get(context.Background())
	if err != nil {
		return err
	}
	var idx models.IndexRef
	err = ref.DataTo(&idx)
	if err != nil {
		return err
	}

	ref, err = db.client.Collection(i.GetType()).Doc(idx.ID).Get(context.Background())
	if err != nil {
		return err
	}
	i.SetID(ref.Ref.ID)
	return ref.DataTo(i)
}

func (db *Store) QueryOne(path, op, value string, i models.Item) error {
	di := db.client.Collection(i.GetType()).Where(path, op, value).Limit(1).Documents(context.Background())
	res, err := di.GetAll()
	if err != nil {
		return err
	}
	if len(res) != 1 {
		return fmt.Errorf("no results found %d", len(res))
	}
	return res[0].DataTo(i)
}

func (db *Store) Query(path, op, value, itemType string) models.Iterator {
	di := db.client.Collection(itemType).Where(path, op, value).Documents(context.Background())
	out := &Iterator{
		i: di,
	}
	return out
}

func (db *Store) All(itemType string) models.Iterator {
	di := db.client.Collection(itemType).Documents(context.Background())
	return &Iterator{
		i: di,
	}
}
