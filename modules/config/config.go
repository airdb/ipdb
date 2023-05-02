package config

type Config struct {
	IPIP *ipdb `mapstructure:"ipip"`
}

type ipdb struct {
	Dataset string `mapstructure:"dataset"`
	Path    string `mapstructure:"path"`
}

var GlobalConfig Config
