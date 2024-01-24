package transformer

import "fmt"

// GetTransformerByName returns a transformer by name
func GetTransformerByName[T any](name string) (Transformer[T], error) {
	switch name {
	case "json-transformer":
		return NewJsonTransformer[T](), nil
	default:
		return nil, fmt.Errorf("transformer %s not found", name)
	}
}
