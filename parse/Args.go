package parse

import "github.com/akyoto/ignore"

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

	for index, char := range command {
		if reader.CanIgnore(char) {
			continue
		}

		if char == ' ' {
			args = append(args, command[argStart:index])
			argStart = index + 1
		}
	}

	// Last argument
	if argStart != len(command)-1 {
		args = append(args, command[argStart:])
	}

	return args
}
