package transformer

import "encoding/json"

type JsonTransformer[T any] struct {
}

func NewJsonTransformer[T any]() *JsonTransformer[T] {
	return &JsonTransformer[T]{}
}

func (t *JsonTransformer[T]) Transform(data []byte) (*T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
