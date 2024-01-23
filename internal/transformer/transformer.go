package transformer

type Transformer[T any] interface {
	Transform(data []byte) (*T, error)
}
