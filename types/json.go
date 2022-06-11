package types

import (
	"encoding/json"
)

type JSONStrategy struct{}

func (j *JSONStrategy) encode(n *Nesting) ([]byte, error) {
	return json.MarshalIndent(n, "", "\t")
}

func (j *JSONStrategy) decode(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, v)
}
