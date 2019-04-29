package main

import (
	"fmt"
	"strconv"
	"strings"
)

func incrementTag(tag string, plus string) (string, error) {
	if !strings.HasPrefix(tag, "v") {
		return "", fmt.Errorf("Not a valid tag, it needs to be prefixed with 'v': %s", tag)
	}

	tag = strings.TrimPrefix(tag, "v")
	parts := strings.Split(tag, ".")
	plusParts := strings.Split(plus, ".")

	if len(parts) < 3 {
		return "", fmt.Errorf("Not a valid SemVer tag: %s", tag)
	}

	for index, part := range parts {
		if index >= len(plusParts) {
			break
		}

		num, err := strconv.Atoi(part)

		if err != nil {
			return "", err
		}

		plusNum, err := strconv.Atoi(plusParts[index])

		if err != nil {
			return "", err
		}

		num += plusNum
		parts[index] = strconv.Itoa(num)
	}

	return "v" + strings.Join(parts, "."), nil
}
