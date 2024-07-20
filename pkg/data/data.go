package data

import "github.com/trenton42/zettel/pkg/types"

type Data struct {
	pth string
}

func New(pth string) (*Data, error) {
	return &Data{
		pth: pth,
	}, nil
}

func (d *Data) Create(_ types.Zettel) error {
	return nil
}

func (d *Data) Update(_ int, _ types.Zettel) error {
	return nil
}

func (d *Data) Get(_ int) (types.Zettel, error) {
	return types.Zettel{}, nil
}

func (d *Data) Delete(_ int) error {
	return nil
}

func (d *Data) Query() ([]types.Zettel, error) {
	return nil, nil
}
