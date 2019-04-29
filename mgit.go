package main

import (
	"flag"
	"path/filepath"
)

var (
	tags         bool
	root         string
	shellCommand string
)

func init() {
	flag.BoolVar(&tags, "tags", false, "Shows the latest tag in every git repository")
	flag.StringVar(&root, "root", ".", "Specifies the directory to search for git repositories")
	flag.StringVar(&shellCommand, "run", "", "Specifies a shell command to execute in every git repository")
	flag.Parse()
}

func main() {
	if !tags && shellCommand == "" {
		flag.Usage()
		return
	}

	err := filepath.Walk(root, walk)

	if err != nil {
		panic(err)
	}

	repositoryWaitGroup.Wait()

	if tags {
		showTags()
	}
}
