package test

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type MockRemote struct {
	ListFunc func(options *git.ListOptions) ([]plumbing.Reference, error)
}

func (m *MockRemote) List(options *git.ListOptions) ([]plumbing.Reference, error) {
	return m.ListFunc(options)
}
