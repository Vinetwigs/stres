package types

import (
	"bytes"

	"github.com/pelletier/go-toml/v2"
)

type TOMLStrategy struct{}

func (t *TOMLStrategy) encode(n *Nesting) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	err := enc.Encode(n)
	return buf.Bytes(), err
}

func (t *TOMLStrategy) decode(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return toml.Unmarshal(data, v)
}
