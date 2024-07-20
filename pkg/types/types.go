package types

type DB interface {
	Create(Zettel) error
	Update(int, Zettel) error
	Get(int) (Zettel, error)
	Delete(int) error
	Query() ([]Zettel, error)
}

type Zettel struct{}

type Index struct {
	Titles map[string][]int `json:"titles"`
	Tags   map[string][]int `json:"tags"`
}
