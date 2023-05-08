package config

type Config struct {
	IPIP        *ipdb `mapstructure:"ipip"`
	IP2Location *ipdb `mapstructure:"ip2location"`
}

type ipdb struct {
	Dataset string `mapstructure:"dataset"`
	Path    string `mapstructure:"path"`
}

var GlobalConfig Config
