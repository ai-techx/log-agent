package config

type Input struct {
	Path string `mapstructure:"path"`
	Name string `mapstructure:"name"`
	Uses string `mapstructure:"uses"`
}

type TransformerSettings struct {
	// ForInput is the name of the input that the transformer is for
	ForInput string `mapstructure:"for_input"`
	// ToOutput is the name of the output that the transformer is for
	ToOutput string `mapstructure:"to_output"`
	// Uses is the name of the transformer to use
	Uses string `mapstructure:"uses"`
}

type ElasticSearch struct {
	Endpoint               string `mapstructure:"endpoint"`
	ApiKey                 string `mapstructure:"api_key"`
	CloudId                string `mapstructure:"cloud_id"`
	ElasticSearchIndexName string `mapstructure:"elastic_search_index_name"`
}

type MongoDB struct {
	// Endpoint is the endpoint of the mongodb server with username and password
	Endpoint       string `mapstructure:"endpoint"`
	CollectionName string `mapstructure:"collection"`
	DatabaseName   string `mapstructure:"database"`
}

type Output struct {
	Name          string         `mapstructure:"name"`
	ElasticSearch *ElasticSearch `mapstructure:"elastic_search"`
	MongoDB       *MongoDB       `mapstructure:"mongodb"`
	Stdout        bool           `mapstructure:"stdout"`
}

type Config struct {
	Input       []Input               `mapstructure:"input"`
	Output      []Output              `mapstructure:"output"`
	Transformer []TransformerSettings `mapstructure:"transformer"`
}
