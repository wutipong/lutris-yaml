package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gosimple/slug"
	"go.yaml.in/yaml/v4"
)

// GameConfig represents the structure of the YAML data.
type GameConfig struct {
	Name     string `yaml:"name"`
	Slug     string `yaml:"slug"`
	GameSlug string `yaml:"game_slug"`
	Version  string `yaml:"version"`
	Runner   string `yaml:"runner"`
	Script   Script `yaml:"script"`
}

// Script represents the game's script configuration.
type Script struct {
	Game GameScript `yaml:"game"`
}

// GameScript represents the game's specific settings.
type GameScript struct {
	Exe string `yaml:"exe"`
	// Prefix string `yaml:"prefix"`
	Arch string `yaml:"arch"`
}

func GenerateInstallFile(name string, path string, config Config) (g GameConfig, err error) {
	slugName := slug.Make(name + " " + filepath.Base(path))

	g = GameConfig{
		Name:     name,
		Slug:     "lutris-yaml-installer",
		GameSlug: slugName,
		Version:  "local",
		Runner:   "wine",

		Script: Script{
			Game: GameScript{
				Exe: path,
				//Prefix: ??
				Arch: "win64",
			},
		},
	}

	b, err := yaml.Marshal(g)
	if err != nil {
		return
	}

	err = os.WriteFile(filepath.Join(config.OutputLocation, fmt.Sprintf("%s.yaml", slugName)), b, 0666)
	return
}
