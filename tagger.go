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
	faintColor          = color.New(color.Faint)
	repositoryWaitGroup sync.WaitGroup
	repositories        []*Repository
)

func main() {
	filepath.Walk(".", walk)
	repositoryWaitGroup.Wait()

	for _, repository := range repositories {
		if repository.LastCommitTagged {
			faintColor.Println(repository.Path)
		} else {
			fmt.Println(repository.Path)
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
	repository.LastCommitHash = getCommitHash(repository.Path)

	if repository.LastCommitHash == "" {
		return
	}

	cmd := exec.Command("git", "describe", "--contains", repository.LastCommitHash)
	cmd.Dir = repository.Path

	_, err := cmd.Output()

	if err != nil {
		return
	}

	repository.LastCommitTagged = true
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
