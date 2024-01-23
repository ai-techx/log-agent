package main

import (
	"github.com/google/logger"
	"io"
	"log-agent/internal/agent"
	"log-agent/internal/output"
	"log-agent/internal/transformer"
	"log-agent/internal/types"
)

func main() {
	logger.Init("Logger", true, false, io.Discard)

	outputClient := output.NewStdout[types.Log]()
	transformerClient := transformer.NewJsonTransformer[types.Log]()

	go func() {
		client := agent.NewTailClient[types.Log]("/Users/sirily11/Desktop/metaverse/log-agent/access.log", outputClient, transformerClient)
		err := client.Read()
		if err != nil {
			logger.Error(err)
		}
	}()

	<-make(chan struct{})
}
