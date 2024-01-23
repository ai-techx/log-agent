package config

type LogFile struct {
	Path string
	Name string
}

type ElasticSearch struct {
	Endpoint               string
	ApiKey                 string
	ElasticSearchIndexName string
}

type Output struct {
	ElasticSearch ElasticSearch
}

type Config struct {
	LogFiles []LogFile
	Output   Output
}
