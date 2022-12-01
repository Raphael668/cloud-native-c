package config

import (
	"cloud-native-c/pkg/db"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server Server    `json:"Server"`
	DB     db.Config `json:"DB"`
	Hello  Hello     `json:"Hello"`
}

type Server struct {
	Port string `json:"Port"`
}

type Hello struct {
	Show string `json:"Show"`
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
