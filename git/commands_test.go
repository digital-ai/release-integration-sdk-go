package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setupTestRepo(t *testing.T) *GitContext {
	ctx, err := InitNewRepository("test-repo")
	if err != nil {
		t.Fatalf("Failed to initialize repository: %v", err)
	}
	return ctx
}

func TestReadFileCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(ctx.repositoryPath+"/"+filePath, content, 0644)
	assert.NoError(t, err)

	cmd := &ReadFileCommand{FilePath: filePath}
	err = ctx.ExecuteCommand(cmd)
	assert.NoError(t, err)
	assert.Equal(t, content, cmd.Content)
}

func TestWriteFileCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")

	cmd := &WriteFileCommand{FilePath: filePath, Content: content}
	err := ctx.ExecuteCommand(cmd)
	assert.NoError(t, err)

	readContent, err := os.ReadFile(ctx.repositoryPath + "/" + filePath)
	assert.NoError(t, err)
	assert.Equal(t, content, readContent)
}

func TestAddFilesCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(ctx.repositoryPath+"/"+filePath, content, 0644)
	assert.NoError(t, err)

	repo := ctx.repo
	w, err := repo.Worktree()
	assert.NoError(t, err)
	status, err := w.Status()
	assert.NoError(t, err)
	assert.Equal(t, git.Untracked, status[filePath].Staging)

	cmd := &AddFilesCommand{Files: []string{filePath}}
	err = ctx.ExecuteCommand(cmd)
	assert.NoError(t, err)

	status, err = w.Status()
	assert.NoError(t, err)
	assert.Equal(t, git.Added, status[filePath].Staging)
}

func TestCommitChangesCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(ctx.repositoryPath+"/"+filePath, content, 0644)
	assert.NoError(t, err)

	addCmd := &AddFilesCommand{Files: []string{filePath}}
	err = ctx.ExecuteCommand(addCmd)
	assert.NoError(t, err)

	commitCmd := &CommitChangesCommand{
		Message:     "Initial commit",
		AuthorName:  "Author Name",
		AuthorEmail: "author@example.com",
	}
	err = ctx.ExecuteCommand(commitCmd)
	assert.NoError(t, err)

	w, err := ctx.repo.Worktree()
	assert.NoError(t, err)
	status, err := w.Status()
	assert.NoError(t, err)
	assert.Empty(t, status)
}

func TestPushChangesCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	auth := &http.BasicAuth{
		Username: "username",
		Password: "password",
	}

	cmd := &PushChangesCommand{
		Auth:   auth,
		Remote: "origin",
		Branch: "main",
	}
	err := ctx.ExecuteCommand(cmd)
	assert.Error(t, err) // Expect error due to no remote repository
}

func TestPullChangesCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	auth := &http.BasicAuth{
		Username: "username",
		Password: "password",
	}

	cmd := &PullChangesCommand{
		Auth:   auth,
		Remote: "origin",
		Branch: "main",
	}
	err := ctx.ExecuteCommand(cmd)
	assert.Error(t, err) // Expect error due to no remote repository
}

func TestCreateBranchCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(ctx.repositoryPath+"/"+filePath, content, 0644)
	assert.NoError(t, err)

	addCmd := &AddFilesCommand{Files: []string{filePath}}
	err = ctx.ExecuteCommand(addCmd)
	assert.NoError(t, err)

	commitCmd := &CommitChangesCommand{
		Message:     "Initial commit",
		AuthorName:  "Author Name",
		AuthorEmail: "author@example.com",
	}
	err = ctx.ExecuteCommand(commitCmd)
	assert.NoError(t, err)

	cmd := &CreateBranchCommand{Branch: "new-branch"}
	err = ctx.ExecuteCommand(cmd)
	assert.NoError(t, err)

	branches, err := ctx.repo.Branches()
	assert.NoError(t, err)
	found := false
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == "new-branch" {
			found = true
		}
		return nil
	})
	assert.NoError(t, err)
	assert.True(t, found)
}

func TestCheckoutBranchCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(ctx.repositoryPath+"/"+filePath, content, 0644)
	assert.NoError(t, err)

	addCmd := &AddFilesCommand{Files: []string{filePath}}
	err = ctx.ExecuteCommand(addCmd)
	assert.NoError(t, err)

	commitCmd := &CommitChangesCommand{
		Message:     "Initial commit",
		AuthorName:  "Author Name",
		AuthorEmail: "author@example.com",
	}
	err = ctx.ExecuteCommand(commitCmd)
	assert.NoError(t, err)

	createCmd := &CreateBranchCommand{Branch: "new-branch"}
	err = ctx.ExecuteCommand(createCmd)
	assert.NoError(t, err)

	cmd := &CheckoutBranchCommand{Branch: "new-branch"}
	err = ctx.ExecuteCommand(cmd)
	assert.NoError(t, err)

	// Verify branch checkout
	head, err := ctx.repo.Head()
	assert.NoError(t, err)
	assert.Equal(t, "refs/heads/new-branch", head.Name().String())
}

func TestMergeBranchCommand(t *testing.T) {
	ctx := setupTestRepo(t)
	defer ctx.Cleanup()

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(ctx.repositoryPath+"/"+filePath, content, 0644)
	assert.NoError(t, err)

	addCmd := &AddFilesCommand{Files: []string{filePath}}
	err = ctx.ExecuteCommand(addCmd)
	assert.NoError(t, err)

	commitCmd := &CommitChangesCommand{
		Message:     "Initial commit",
		AuthorName:  "Author Name",
		AuthorEmail: "author@example.com",
	}
	err = ctx.ExecuteCommand(commitCmd)
	assert.NoError(t, err)

	createCmd := &CreateBranchCommand{Branch: "source-branch"}
	err = ctx.ExecuteCommand(createCmd)
	assert.NoError(t, err)

	createCmd = &CreateBranchCommand{Branch: "target-branch"}
	err = ctx.ExecuteCommand(createCmd)
	assert.NoError(t, err)

	cmd := &MergeBranchCommand{SourceBranch: "source-branch", TargetBranch: "target-branch"}
	err = ctx.ExecuteCommand(cmd)
	assert.Error(t, err)
}
