package dagcampaign

import (
	"errors"
	"io"

	"github.com/goccy/go-graphviz"
)

type Campaign struct {
	Name string
}

func (c *Campaign) Render(format graphviz.Format, w io.Writer) error {
	return errors.New("no implemented")
}
