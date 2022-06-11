package types

import (
	yaml "gopkg.in/yaml.v3"
)

type YAMLStrategy struct{}

func (y *YAMLStrategy) encode(n *Nesting) ([]byte, error) {
	return yaml.Marshal(n)
}

func (y *YAMLStrategy) decode(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return yaml.Unmarshal(data, v)
}
