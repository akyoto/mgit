package parse_test

import (
	"github.com/akyoto/mgit/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleParse(t *testing.T) {
	args := parse.Args("git status")
	assert.Len(t, args, 2)
	assert.Equal(t, args[0], "git")
	assert.Equal(t, args[1], "status")
}

func TestAdvancedParse(t *testing.T) {
	args := parse.Args("git commit -m \"My update message\"")
	assert.Len(t, args, 4)
	assert.Equal(t, args[0], "git")
	assert.Equal(t, args[1], "commit")
	assert.Equal(t, args[2], "-m")
	assert.Equal(t, args[3], "\"My update message\"")
}
