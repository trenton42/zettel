package data

import (
	"os"
	"time"

	"github.com/trenton42/zettel/pkg/types"
)

type Data struct {
	pth   string
	index *types.Index
}

func New(pth string) (*Data, error) {
	_, err := os.Stat(pth)
	if err != nil {
		err = os.Mkdir(pth, 0744)
		if err != nil {
			return nil, err
		}
	}
	d := &Data{
		pth: pth,
	}
	return d, d.loadIndex()
}

func (d *Data) Create(z types.Zettel) error {
	z.Created = time.Now()
	z.Updated = z.Created
	z.ID = d.NextIndex()
	d.saveZettel(z)
	return nil
}

func (d *Data) Update(id int, z types.Zettel) error {
	old, err := d.loadZettel(id)
	if err != nil {
		return err
	}
	old.Body = z.Body
	old.Title = z.Title
	old.Updated = time.Now()
	d.saveZettel(old)
	return nil
}

func (d *Data) Get(id int) (types.Zettel, error) {
	return d.loadZettel(id)
}

func (d *Data) Delete(id int) error {
	return d.deleteZettel(id)
}

func (d *Data) Query() ([]types.Zettel, error) {
	return nil, nil
}

func (d *Data) List() (map[int]string, error) {
	return d.index.IDs, nil
}
