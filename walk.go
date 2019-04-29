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
	"github.com/akyoto/mgit/parse"
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
		repository.VisualizedTag = repository.LastTag

		if len(repository.VisualizedTag) > maxTagLength {
			maxTagLength = len(repository.VisualizedTag)
		}

		if repository.LastCommitHash == "" {
			return
		}

		repository.LastCommitTagged = isLastCommitTagged(repository.Path, repository.LastCommitHash)
	}

	if tag != "" && repository.LastTag != "" && !repository.LastCommitTagged {
		newTag, err := incrementTag(repository.LastTag, tag[1:])

		if err != nil {
			return
		}

		if setTag(repository, newTag) != nil {
			return
		}

		repository.NewTag = newTag
		repository.VisualizedTag = fmt.Sprintf("%s -> %s", repository.LastTag, repository.NewTag)

		if len(repository.VisualizedTag) > maxTagLength {
			maxTagLength = len(repository.VisualizedTag)
		}
	}
}

func setTag(repository *Repository, newTag string) error {
	if dry {
		return nil
	}

	cmd := exec.Command("git", "tag", newTag)
	cmd.Dir = repository.Path
	_, err := cmd.Output()

	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push", "origin", newTag)
	cmd.Dir = repository.Path
	_, err = cmd.Output()
	return err
}

func runCommandInDirectory(command string, repository *Repository) {
	fmt.Printf("\033[2K\rRunning %s in %s", color.GreenString(command), color.GreenString(repository.Path))

	var cmd *exec.Cmd
	args := parse.Args(command)

	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0])
	}

	cmd.Dir = repository.Path
	output, err := cmd.CombinedOutput()
	repository.Command.Output = strings.TrimSpace(string(output))
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
