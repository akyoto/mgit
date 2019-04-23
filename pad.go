package main

import "strings"

func rightPad(s string, pad string, length int) string {
	return s + strings.Repeat(pad, length-len(s))
}
