package types

import "github.com/genkami/watson"

type WatsonStrategy struct{}

func (w *WatsonStrategy) encode(n *Nesting) ([]byte, error) {
	return watson.Marshal(n)
}

func (w *WatsonStrategy) decode(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return watson.Unmarshal(data, v)
}
