package agent

type Client[T any] interface {
	Read() error
}
