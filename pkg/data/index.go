package data

import (
	"encoding/json"
	"os"
	"path"

	"github.com/trenton42/zettel/pkg/types"
)

func (d *Data) indexPath() string {
	return path.Join(d.pth, "index.json")
}

func (d *Data) loadIndex() error {
	fp, err := os.Open(d.indexPath())
	if err != nil {
		d.index = &types.Index{
			Titles: make(map[string][]int),
			Tags:   make(map[string][]int),
			IDs:    make(map[int]string),
			Next:   1,
		}
		return nil
	}
	defer fp.Close()
	d.index = &types.Index{}
	dec := json.NewDecoder(fp)
	return dec.Decode(d.index)
}

func (d *Data) saveIndex() error {
	fp, err := os.OpenFile(d.indexPath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	enc := json.NewEncoder(fp)
	return enc.Encode(d.index)
}

func (d *Data) updateIndex(z types.Zettel, rm bool) {
	if rm {
		// Deleting a zettel
		delete(d.index.IDs, z.ID)
		// TODO update the rest
		return
	}
	d.index.IDs[z.ID] = z.Title
	d.index.Titles[z.Title] = append(d.index.Titles[z.Title], z.ID)
}

func (d *Data) NextIndex() (out int) {
	out = d.index.Next
	d.index.Next++
	return
}
