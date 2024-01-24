package agent

import (
	"log-agent/internal/output"
	"log-agent/internal/transformer"
)

type Client[T any] interface {
	Read() error
	SetTransformer(transformer transformer.Transformer[T])
	SetOutput(output output.Client[T])
}
