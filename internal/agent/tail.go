package agent

import (
	"github.com/google/logger"
	"github.com/nxadm/tail"
	"io"
	"log-agent/internal/output"
	"log-agent/internal/transformer"
)

type TailClient[T any] struct {
	filePath     string
	outputClient output.Client[T]
	transformer  transformer.Transformer[T]
}

// NewTailClient creates a new tail client
// This client will continuously tail a file and write the outputClient to the outputClient client
func NewTailClient[T any](path string, outputClient output.Client[T], transformer transformer.Transformer[T]) *TailClient[T] {
	return &TailClient[T]{
		filePath:     path,
		outputClient: outputClient,
		transformer:  transformer,
	}
}

func (c *TailClient[T]) Read() error {
	logger.Info("Using Tail Client: ", c.filePath)
	config := tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
		Logger: tail.DiscardingLogger,
	}
	t, err := tail.TailFile(
		c.filePath, config)
	if err != nil {
		return err
	}

	// Print the text of each received line
	for line := range t.Lines {
		transformed, err := c.transformer.Transform([]byte(line.Text))
		if err != nil {
			return err
		}

		err = c.outputClient.Write(*transformed)
		if err != nil {
			return err
		}
	}

	return nil
}
