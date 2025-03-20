package git

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"path/filepath"
)

type ReadFileCommand struct {
	FilePath string
	Content  []byte
}

func (c *ReadFileCommand) execute(repo *git.Repository) error {
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	absPath := filepath.Join(worktree.Filesystem.Root(), c.FilePath)
	content, err := os.ReadFile(absPath)
	if err != nil {
		return err
	}

	c.Content = content
	return nil
}

type WriteFileCommand struct {
	FilePath string
	Content  []byte
}

func (c *WriteFileCommand) execute(repo *git.Repository) error {
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	absPath := filepath.Join(worktree.Filesystem.Root(), c.FilePath)
	err = os.WriteFile(absPath, c.Content, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

type AddFilesCommand struct {
	Files []string
}

func (c *AddFilesCommand) execute(repo *git.Repository) error {
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	for _, file := range c.Files {
		_, err = w.Add(file)
		if err != nil {
			return err
		}
	}
	return nil
}

type CommitChangesCommand struct {
	Message string
}

func (c *CommitChangesCommand) execute(repo *git.Repository) error {
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Commit(c.Message, &git.CommitOptions{})
	return err
}

type PushChangesCommand struct {
	Auth   *http.BasicAuth
	Remote string
	Branch string
}

func (c *PushChangesCommand) execute(repo *git.Repository) error {
	return repo.Push(&git.PushOptions{
		RemoteName: c.Remote,
		Auth:       c.Auth,
	})
}

type PullChangesCommand struct {
	Auth   *http.BasicAuth
	Remote string
	Branch string
}

func (c *PullChangesCommand) execute(repo *git.Repository) error {
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	return w.Pull(&git.PullOptions{
		RemoteName: c.Remote,
		Auth:       c.Auth,
	})
}

type CreateBranchCommand struct {
	Branch string
}

func (c *CreateBranchCommand) execute(repo *git.Repository) error {
	headRef, err := repo.Head()
	if err == plumbing.ErrReferenceNotFound {
		// initial commit if the repository is empty
		return errors.New("initial commit not in place - can't create branch")
	} else if err != nil {
		return err
	}

	refName := plumbing.NewBranchReferenceName(c.Branch)
	ref := plumbing.NewHashReference(refName, headRef.Hash())

	return repo.Storer.SetReference(ref)
}

type CheckoutBranchCommand struct {
	Branch string
}

func (c *CheckoutBranchCommand) execute(repo *git.Repository) error {
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	return w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(c.Branch),
	})
}

type MergeBranchCommand struct {
	SourceBranch string
	TargetBranch string
}

func (c *MergeBranchCommand) execute(repo *git.Repository) error {
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(c.TargetBranch),
	})
	if err != nil {
		return err
	}

	sourceRef, err := repo.Reference(plumbing.NewBranchReferenceName(c.SourceBranch), true)
	if err != nil {
		return err
	}

	headRef, err := repo.Head()
	if err != nil {
		return err
	}

	mergeCommit, err := w.Commit("Merge branch "+c.SourceBranch+" into "+c.TargetBranch, &git.CommitOptions{
		Parents: []plumbing.Hash{headRef.Hash(), sourceRef.Hash()},
	})
	if err != nil {
		return err
	}

	targetRef := plumbing.NewHashReference(plumbing.NewBranchReferenceName(c.TargetBranch), mergeCommit)
	return repo.Storer.SetReference(targetRef)
}
