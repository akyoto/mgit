package parse_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/mgit/parse"
)

func TestSimpleArgs(t *testing.T) {
	args := parse.Args("git status")
	assert.Equal(t, len(args), 2)
	assert.Equal(t, args[0], "git")
	assert.Equal(t, args[1], "status")
}

func TestAdvancedArgs(t *testing.T) {
	args := parse.Args("git commit -m \"My update message\"")
	assert.Equal(t, len(args), 4)
	assert.Equal(t, args[0], "git")
	assert.Equal(t, args[1], "commit")
	assert.Equal(t, args[2], "-m")
	assert.Equal(t, args[3], "My update message")
}

func TestWeirdArgs(t *testing.T) {
	args := parse.Args("\"git\" \"\" \"oh\" \"oh noes\" \"this must not break\"")
	assert.Equal(t, len(args), 5)
	assert.Equal(t, args[0], "git")
	assert.Equal(t, args[1], "")
	assert.Equal(t, args[2], "oh")
	assert.Equal(t, args[3], "oh noes")
	assert.Equal(t, args[4], "this must not break")
}
