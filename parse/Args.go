package parse

import (
	"github.com/akyoto/ignore"
	"strings"
)

// Args parses arguments and makes sure that string arguments
// are correctly counted as a single argument.
//
// Example:
//
//     git commit -m "Testing splitting"
//
// A naive strings.Split(" ") would result in 5 arguments:
//
//     [git, commit, -m, "Testing, splitting"]
//
// This function instead generates 4 arguments:
//
//     [git, commit, -m, "Testing splitting"]
//
func Args(command string) []string {
	var args []string
	reader := ignore.Reader{}
	argStart := 0

	if !strings.HasSuffix(command, " ") {
		command += " "
	}

	for index, char := range command {
		if reader.CanIgnore(char) {
			continue
		}

		if char == ' ' {
			offset := 0

			if (command[argStart] == '"' && command[index-1] == '"') || (command[argStart] == '\'' && command[index-1] == '\'') {
				offset++
			}

			args = append(args, command[argStart+offset:index-offset])
			argStart = index + 1
		}
	}

	return args
}
