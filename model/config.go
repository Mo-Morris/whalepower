package model

type Config struct {
	Port int `json:"port,omitempty" mapstructure:"port" yaml:"port"`
}
