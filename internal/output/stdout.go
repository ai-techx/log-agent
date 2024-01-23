package output

import "github.com/google/logger"

type Stdout[T any] struct {
}

func NewStdout[T any]() *Stdout[T] {
	return &Stdout[T]{}
}

func (s *Stdout[T]) Write(data T) error {
	logger.Info(data)
	return nil
}
