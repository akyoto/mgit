package main

// Repository represents a git repository.
type Repository struct {
	Path             string
	LastCommitHash   string
	LastCommitTagged bool
	LastTag          string
	Command          struct {
		Output string
		Error  error
	}
}
