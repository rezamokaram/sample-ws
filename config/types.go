package config

type ServerConfig struct {
	HttpPort uint   `json:"http_port" yaml:"http_port"`
	HttpHost string `json:"http_host" yaml:"http_host"`
	Name     string `json:"name" yaml:"name"`
	Version  string `json:"version" yaml:"version"`
}
