package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server Server `json:"Server"`
}

type Server struct {
	Port string `json:"Port"`
}

func New(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
