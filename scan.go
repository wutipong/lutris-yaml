package main

import (
	"path/filepath"
	"slices"
)

func PopulateExecutables(config Config) (paths []string, err error) {
	src, err := filepath.Abs(config.GamesLocation)
	if err != nil {
		return
	}

	pattern := filepath.Join(src, "**", "*.exe")

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return
	}

	for _, path := range matches {
		filename := filepath.Base(path)

		if slices.Contains(config.IgnoredFiles, filename) {
			continue
		}

		paths = append(paths, path)
	}

	return
}
