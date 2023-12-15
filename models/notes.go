package models

type Zettel struct {
	IDField
	Title string `json:"title"`
	Body  string `json:"body"`
	Owner string `json:"-"`
}

func (z *Zettel) GetType() string {
	return "zettles"
}
