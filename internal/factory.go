package internal

import "github.com/cli/go-gh/v2/pkg/repository"

type Factory struct {
	RepositoryResolver func() (repository.Repository, error)
}
