package main

import (
	"github.com/akyoto/color"
)

var (
	maxRepositoryLength = 0
	maxTagLength        = 0
)

func showTags() {
	for _, repository := range repositories {
		pathPadded := rightPad(repository.Path, " ", maxRepositoryLength)
		tagPadded := rightPad(repository.LastTag, " ", maxTagLength)

		if repository.LastCommitTagged {
			color.Green("%s | %s", pathPadded, tagPadded)
		} else {
			if repository.LastTag == "" {
				color.Red("%s | %s | not tagged", pathPadded, tagPadded)
			} else {
				color.Yellow("%s | %s | outdated", pathPadded, tagPadded)
			}
		}
	}
}
