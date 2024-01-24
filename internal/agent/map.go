package agent

import (
	"fmt"
	"log-agent/internal/config"
)

func GetAgentByConfig[T any](conf config.Input) (Client[T], error) {
	switch conf.Uses {
	case "tail":
		return NewTailClient[T](conf.Path), nil
	default:
		return nil, fmt.Errorf("agent %s not found", conf.Uses)
	}
}
