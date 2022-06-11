package types

type StrategyAlgo interface {
	decode(data []byte, v interface{}) error
	encode(n *Nesting) ([]byte, error)
}

type EncoderDecoder struct {
	strategy StrategyAlgo
}

func (e *EncoderDecoder) SetStrategy(s StrategyAlgo) {
	e.strategy = s
}

func (e *EncoderDecoder) GetStrategy() StrategyAlgo {
	return e.strategy
}

func (e *EncoderDecoder) Encode(n *Nesting) ([]byte, error) {
	return e.strategy.encode(n)
}

func (e *EncoderDecoder) Decode(data []byte, v interface{}) error {
	return e.strategy.decode(data, v)
}
