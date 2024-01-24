package main

import (
	"github.com/google/logger"
	"github.com/spf13/viper"
	"io"
	"log-agent/internal/agent"
	"log-agent/internal/config"
	"log-agent/internal/output"
	"log-agent/internal/transformer"
)

// readConfig reads the config file and returns the config using viper
func readConfigFromFile() (*config.Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.AddConfigPath("/etc/cron-requests/")
	viper.AddConfigPath("$HOME/.cron-requests")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var readConfig config.Config
	err = viper.Unmarshal(&readConfig)
	if err != nil {
		return nil, err
	}

	return &readConfig, nil
}

func main() {
	logger.Init("Logger", true, false, io.Discard)

	readConfig, err := readConfigFromFile()
	if err != nil {
		panic(err)
	}

	inputClients := make(map[string]agent.Client[any])
	outputClients := make(map[string]output.Client[any])
	transformers := make(map[string]transformer.Transformer[any])

	for _, t := range readConfig.Transformer {
		transformerClient, err := transformer.GetTransformerByName[any](t.Uses)
		if err != nil {
			panic(err)
		}
		transformers[t.Uses] = transformerClient
	}

	for _, o := range readConfig.Output {
		outputClient, err := output.GetOutputByConfig[any](o)
		if err != nil {
			panic(err)
		}
		outputClients[o.Name] = outputClient
	}

	for _, i := range readConfig.Input {
		inputClient, err := agent.GetAgentByConfig[any](i)
		if err != nil {
			panic(err)
		}
		inputClients[i.Name] = inputClient
	}

	for _, t := range readConfig.Transformer {
		inputClient := setupInputClient(t, transformers, inputClients, outputClients)
		t := t
		go func() {
			logger.Infof("Starting transforming input %s to output %s using %s", t.ForInput, t.ToOutput, t.Uses)
			err := inputClient.Read()
			if err != nil {
				logger.Errorf("Error reading input %s: %s", t.ForInput, err)
			}
		}()
	}

	<-make(chan struct{})
}

func setupInputClient(t config.TransformerSettings, transformers map[string]transformer.Transformer[any], inputClients map[string]agent.Client[any], outputClients map[string]output.Client[any]) agent.Client[any] {
	transformerClient := transformers[t.Uses]
	inputClient := inputClients[t.ForInput]
	outputClient := outputClients[t.ToOutput]

	if inputClient == nil {
		logger.Fatal("Input client not found: ", t.ForInput)
	}

	if outputClient == nil {
		logger.Fatal("Output client not found: ", t.ToOutput)
	}
	inputClient.SetTransformer(transformerClient)
	inputClient.SetOutput(outputClient)
	return inputClient
}
