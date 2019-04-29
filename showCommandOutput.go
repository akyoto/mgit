package main

import (
	"fmt"
	"github.com/akyoto/color"
)

func showCommandOutput() {
	for _, repository := range repositories {
		out := color.Green

		if repository.Command.Error != nil {
			out = color.Red
		}

		out(repository.Path + ":")
		fmt.Println(repository.Command.Output)
	}
}
