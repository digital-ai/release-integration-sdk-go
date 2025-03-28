package git

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"path/filepath"
	"time"
)

// ReadFileCommand is a command that reads a file from a git repository.
type ReadFileCommand struct {
	FilePath string
	Content  []byte
}

func (c *ReadFileCommand) execute(ctx *GitContext) error {
	worktree, err := ctx.repo.Worktree()
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

// WriteFileCommand is a command that writes a file to a git repository.
type WriteFileCommand struct {
	FilePath string
	Content  []byte
}

func (c *WriteFileCommand) execute(ctx *GitContext) error {
	worktree, err := ctx.repo.Worktree()
	if err != nil {
		return err
	}

	absPath := filepath.Join(worktree.Filesystem.Root(), c.FilePath)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(absPath), os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = os.WriteFile(absPath, c.Content, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFileCommand is a command that removes a file from a git repository.
type RemoveFileCommand struct {
	FilePath string
}

func (c *RemoveFileCommand) execute(ctx *GitContext) error {
	worktree, err := ctx.repo.Worktree()
	if err != nil {
		return err
	}

	absPath := filepath.Join(worktree.Filesystem.Root(), c.FilePath)
	err = os.Remove(absPath)
	if err != nil {
		return err
	}

	return nil
}

// AddFilesCommand is a command that adds files to the git index.
type AddFilesCommand struct {
	Files []string
}

func (c *AddFilesCommand) execute(ctx *GitContext) error {
	w, err := ctx.repo.Worktree()
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

// CommitChangesCommand is a command that commits changes to the git repository.
type CommitChangesCommand struct {
	Message     string
	AuthorName  string
	AuthorEmail string
}

func (c *CommitChangesCommand) execute(ctx *GitContext) error {
	w, err := ctx.repo.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Commit(c.Message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  c.AuthorName,
			Email: c.AuthorEmail,
			When:  time.Now(),
		},
	})
	return err
}

// PushChangesCommand is a command that pushes changes to a remote git repository.
type PushChangesCommand struct {
	Remote string
	Branch string
}

func (c *PushChangesCommand) execute(ctx *GitContext) error {
	return ctx.repo.Push(&git.PushOptions{
		RemoteName: c.Remote,
		Auth:       ctx.authMethod,
	})
}

// PullChangesCommand is a command that pulls changes from a remote git repository.
type PullChangesCommand struct {
	Remote string
	Branch string
}

func (c *PullChangesCommand) execute(ctx *GitContext) error {
	w, err := ctx.repo.Worktree()
	if err != nil {
		return err
	}
	return w.Pull(&git.PullOptions{
		RemoteName: c.Remote,
		Auth:       ctx.authMethod,
	})
}

// CreateBranchCommand is a command that creates a new branch in the git repository.
type CreateBranchCommand struct {
	Branch string
}

func (c *CreateBranchCommand) execute(ctx *GitContext) error {
	headRef, err := ctx.repo.Head()
	if err == plumbing.ErrReferenceNotFound {
		// initial commit if the repository is empty
		return errors.New("initial commit not in place - can't create branch")
	} else if err != nil {
		return err
	}

	refName := plumbing.NewBranchReferenceName(c.Branch)
	ref := plumbing.NewHashReference(refName, headRef.Hash())

	return ctx.repo.Storer.SetReference(ref)
}

// CheckoutBranchCommand is a command that checks out a branch in the git repository.
type CheckoutBranchCommand struct {
	Branch string
}

func (c *CheckoutBranchCommand) execute(ctx *GitContext) error {
	w, err := ctx.repo.Worktree()
	if err != nil {
		return err
	}
	return w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(c.Branch),
	})
}

// MergeBranchCommand is a command that merges a branch into another branch in the git repository.
type MergeBranchCommand struct {
	SourceBranch string
	TargetBranch string
}

func (c *MergeBranchCommand) execute(ctx *GitContext) error {
	w, err := ctx.repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(c.TargetBranch),
	})
	if err != nil {
		return err
	}

	sourceRef, err := ctx.repo.Reference(plumbing.NewBranchReferenceName(c.SourceBranch), true)
	if err != nil {
		return err
	}

	headRef, err := ctx.repo.Head()
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
	return ctx.repo.Storer.SetReference(targetRef)
}
