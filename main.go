package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	paths, err := PopulateExecutables(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll(config.OutputLocation, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range paths {
		name, err := GetName(p, config)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = GenerateInstallFile(name, p, config)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
