package main

import (
	"io"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	GamesLocation  string   `yaml:"games_location"`
	OutputLocation string   `yaml:"output_location"`
	IgnoredFiles   []string `yaml:"ignored_files"`
}

func GetConfig() (config Config, err error) {
	config = Config{
		GamesLocation:  "./games",
		OutputLocation: "./output",
		IgnoredFiles:   make([]string, 0),
	}

	f, err := os.Open("./config.yaml")
	if err == nil {
		b, e := io.ReadAll(f)
		if e != nil {
			err = e
			return
		}

		err = yaml.Unmarshal(b, &config)
		if err != nil {
			return
		}
	}

	return
}
