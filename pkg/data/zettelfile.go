package data

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/adrg/frontmatter"
	"github.com/trenton42/zettel/pkg/types"
	"gopkg.in/yaml.v2"
)

func (d *Data) saveZettel(z types.Zettel) error {
	fp, err := os.OpenFile(d.zettelPath(z.ID), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	fp.Write([]byte("---\n"))
	enc := yaml.NewEncoder(fp)
	err = enc.Encode(z)
	fp.Write([]byte("\n---\n"))
	fp.Write([]byte(z.Body))
	d.updateIndex(z, false)
	d.saveIndex()
	return err
}

func (d *Data) loadZettel(id int) (z types.Zettel, err error) {
	var fp io.Reader
	var rest []byte
	fp, err = os.Open(d.zettelPath(id))
	if err != nil {
		return
	}
	rest, err = frontmatter.Parse(fp, &z)
	if err != nil {
		return
	}
	z.Body = string(rest)
	return
}

func (d *Data) deleteZettel(id int) error {
	d.updateIndex(types.Zettel{ID: id}, true)
	d.saveIndex()
	return os.Remove(d.zettelPath(id))
}

func (d *Data) zettelPath(id int) string {
	return path.Join(d.pth, fmt.Sprintf("%d.md", id))
}
