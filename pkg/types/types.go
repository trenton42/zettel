package types

import "time"

type DB interface {
	Create(Zettel) error
	Update(int, Zettel) error
	Get(int) (Zettel, error)
	Delete(int) error
	Query() ([]Zettel, error)
	List() (map[int]string, error)
}

type Zettel struct {
	ID      int       `yaml:"id"`
	Title   string    `yaml:"title"`
	Created time.Time `yaml:"created"`
	Updated time.Time `yaml:"updated"`
	Body    string    `yaml:"-"`
}

type Index struct {
	Titles map[string][]int `json:"titles"`
	Tags   map[string][]int `json:"tags"`
	IDs    map[int]string   `json:"ids"`
	Next   int              `json:"next"`
}
