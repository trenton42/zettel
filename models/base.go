package models

import "time"

type IDField struct {
	ID string `firestore:"-" json:"-"`
}

func (id *IDField) GetID() string {
	return id.ID
}

func (id *IDField) SetID(val string) {
	id.ID = val
}

type TTLField struct {
	TTL time.Time `firestore:"ttl" json:"expires"`
}

func (t *TTLField) SetTTL(val time.Time) {
	t.TTL = val
}

type IndexRef struct {
	ID   string `firestore:"id"`
	Type string `firestore:"type"`
}

type Item interface {
	GetID() string
	SetID(string)
	GetType() string
}

type TTL interface {
	SetTTL(time.Time)
}

type Indexed interface {
	Item
	GetIndex() string
}

type Iterator interface {
	Next(Item) error
}
