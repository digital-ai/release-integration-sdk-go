package test

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/mock"
)

type MockRemote struct {
	mock.Mock
}

func (m *MockRemote) List(o *git.ListOptions) ([]*plumbing.Reference, error) {
	args := m.Called(o)
	return args.Get(0).([]*plumbing.Reference), args.Error(1)
}
