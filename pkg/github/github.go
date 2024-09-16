package github

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
)

type Github struct {
	repository *git.Repository
}

func NewGithub() *Github {
	repo, err := cloneRepo()
	if err != nil {
		//TODO: Handle error
	}
	return &Github{repository: repo}
}

func cloneRepo() (repository *git.Repository, err error) {
	fs := memfs.New()
	storage := memory.NewStorage()
	repo, err := git.Clone(storage, fs, &git.CloneOptions{
		URL:      "", //TODO: Add URL Const
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (g *Github) GetChangedProjects(oldVersion string, newVersion string) ([]string, error) {
	var result []string

	return result, nil
}
