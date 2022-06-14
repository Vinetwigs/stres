package types

import "github.com/vmihailenco/msgpack/v5"

type MsgPackStrategy struct{}

func (m *MsgPackStrategy) encode(n *Nesting) ([]byte, error) {
	return msgpack.Marshal(n)
}

func (m *MsgPackStrategy) decode(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return msgpack.Unmarshal(data, v)
}
