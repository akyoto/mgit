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
	repositoryWaitGroup sync.WaitGroup
	repositories        []*Repository
)

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
	if shellCommand != "" {
		runCommandInDirectory(shellCommand, repository)
	}

	if tags {
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
}

func runCommandInDirectory(command string, repository *Repository) {
	fmt.Printf("\033[2K\rRunning %s in %s", color.GreenString(command), color.GreenString(repository.Path))

	var cmd *exec.Cmd
	args := strings.Split(command, " ")

	if len(args) > 0 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(command)
	}

	cmd.Dir = repository.Path
	output, err := cmd.CombinedOutput()
	repository.Command.Output = string(output)
	repository.Command.Error = err
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
