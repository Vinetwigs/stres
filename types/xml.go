package types

import "encoding/xml"

type XMLStrategy struct{}

func (x *XMLStrategy) encode(n *Nesting) ([]byte, error) {
	return xml.MarshalIndent(n, "", "\t")
}

func (x *XMLStrategy) decode(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return xml.Unmarshal(data, v)
}
