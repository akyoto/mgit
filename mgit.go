package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	tags                 bool
	dry                  bool
	tag                  string
	root                 string
	shellCommand         string
	excludedRepositories string
	excluded             map[string]bool
)

func init() {
	flag.BoolVar(&tags, "tags", false, "Shows the latest tag in every git repository")
	flag.BoolVar(&dry, "dry", false, "Disables the actual tagging and only shows the tag diffs (only applicable with -tag option)")
	flag.StringVar(&tag, "tag", "", "Specifies the tag increment for every outdated git repository (syntax: +0.0.1 to increase the SemVer minor)")
	flag.StringVar(&root, "root", ".", "Specifies the directory to search for git repositories")
	flag.StringVar(&shellCommand, "run", "", "Specifies a shell command to execute in every git repository")
	flag.StringVar(&excludedRepositories, "exclude", "", "Comma separated list of repositories to exclude")
	flag.Parse()
}

func main() {
	// Show help
	if !tags && tag == "" && shellCommand == "" {
		flag.Usage()
		return
	}

	// Tag activates "show tags"
	if tag != "" {
		tags = true

		if !strings.HasPrefix(tag, "+") {
			fmt.Println("Syntax: mgit -tag +0.0.1")
			return
		}
	}

	if dry && tag == "" {
		fmt.Println("Dry run is only possible when -tag option is specified")
		return
	}

	if excludedRepositories != "" {
		excludedList := strings.Split(excludedRepositories, ",")
		excluded = map[string]bool{}

		for _, repo := range excludedList {
			excluded[repo] = true
		}
	}

	// Process repositories in parallel
	err := filepath.Walk(root, walk)

	if err != nil {
		panic(err)
	}

	repositoryWaitGroup.Wait()

	// Shell command output
	if shellCommand != "" {
		fmt.Print("\033[2K\r")
		showCommandOutput()
	}

	// Tags output
	if tags {
		if shellCommand != "" {
			fmt.Println()
		}

		showTags()
	}
}
