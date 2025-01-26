package config

import "github.com/BurntSushi/toml"

type Config struct {
	BranchNamePattern string `toml:"branch_name_pattern"`
}

func NewConfigFromFile(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
