package main

import (
	"flag"
	"path/filepath"
)

var (
	run string
)

func init() {
	flag.StringVar(&run, "run", "", "Specifies a shell command to execute in every git repository.")
	flag.Parse()
}

func main() {
	err := filepath.Walk(".", walk)

	if err != nil {
		panic(err)
	}

	repositoryWaitGroup.Wait()
	showTags()
}
