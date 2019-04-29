package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/akyoto/color"
)

var (
	maxRepositoryLength = 0
	maxTagLength        = 0
	repositoryWaitGroup sync.WaitGroup
	repositories        []*Repository
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

func walk(file string, info os.FileInfo, err error) error {
	if info == nil || !info.IsDir() {
		return nil
	}

	name := info.Name()

	if name != "." && strings.HasPrefix(name, ".") {
		if name == ".git" {
			repositoryPath := strings.TrimSuffix(file, ".git")
			repositoryPath = path.Clean(repositoryPath)

			if len(repositoryPath) > maxRepositoryLength {
				maxRepositoryLength = len(repositoryPath)
			}

			repository := &Repository{
				Path: repositoryPath,
			}

			repositoryWaitGroup.Add(1)

			go func() {
				processRepository(repository)
				repositoryWaitGroup.Done()
			}()

			repositories = append(repositories, repository)
		}

		return filepath.SkipDir
	}

	return nil
}

func processRepository(repository *Repository) {
	if run != "" {
		runUserCommand(repository.Path)
	}

	repository.LastCommitHash = getCommitHash(repository.Path)
	repository.LastTag = getLatestTag(repository.Path)

	if len(repository.LastTag) > maxTagLength {
		maxTagLength = len(repository.LastTag)
	}

	if repository.LastCommitHash == "" {
		return
	}

	repository.LastCommitTagged = isLastCommitTagged(repository.Path, repository.LastCommitHash)
}

func runUserCommand(dir string) {
	fmt.Printf("Running %s in %s\n", color.GreenString(run), color.GreenString(dir))

	var cmd *exec.Cmd
	args := strings.Split(run, " ")

	if len(args) > 0 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(run)
	}

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	_ = cmd.Run()
}

func isLastCommitTagged(dir string, commitHash string) bool {
	cmd := exec.Command("git", "describe", "--contains", commitHash)
	cmd.Dir = dir
	_, err := cmd.Output()

	if err != nil {
		return false
	}

	return true
}

func getCommitHash(dir string) string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = dir
	out, err := cmd.Output()

	if err != nil {
		color.Red(err.Error())
		return ""
	}

	out = bytes.TrimSpace(out)
	return string(out)
}

func getLatestTag(dir string) string {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	cmd.Dir = dir
	out, err := cmd.Output()

	if err != nil {
		return ""
	}

	out = bytes.TrimSpace(out)
	return string(out)
}
