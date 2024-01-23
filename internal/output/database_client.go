package output

type Client[T any] interface {
	Write(data T) error
}
