package output

import (
	"fmt"
	"log-agent/internal/config"
)

func GetOutputByConfig[T any](conf config.Output) (Client[T], error) {
	if conf.ElasticSearch != nil {
		return NewElasticSearchClient[T](*conf.ElasticSearch), nil
	}

	if conf.Stdout == true {
		return NewStdout[T](), nil
	}
	return nil, fmt.Errorf("output client not found")
}
