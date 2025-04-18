package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	auth := &http.BasicAuth{
		Username: "username",
		Password: "password",
	}
	repoURL := "https://example.com/repo.git"
	ctx := NewRepository(repoURL).WithBasicAuth(auth.Username, auth.Password)

	assert.Equal(t, filepath.Base(repoURL), filepath.Base(ctx.repositoryPath))
	assert.Equal(t, repoURL, ctx.repoURL)
	assert.Equal(t, auth, ctx.authMethod)
}

func TestInitNewRepository(t *testing.T) {
	repoName := "test-repo"
	ctx, err := InitNewRepository(repoName)
	defer ctx.Cleanup()

	assert.NoError(t, err)
	assert.NotNil(t, ctx.repo)
	assert.Equal(t, filepath.Join(os.TempDir(), repoName), ctx.repositoryPath)
}

func TestGitContext_WithProxy(t *testing.T) {
	ctx := &GitContext{}
	proxyURL := "http://proxy.example.com"
	proxyUsername := "user"
	proxyPassword := "pass"

	ctx = ctx.WithProxy(proxyURL, proxyUsername, proxyPassword)

	assert.Equal(t, proxyURL, ctx.proxyURL)
	assert.Equal(t, proxyUsername, ctx.proxyUsername)
	assert.Equal(t, proxyPassword, ctx.proxyPassword)
}

func TestGitContext_WithReferenceName(t *testing.T) {
	ctx := &GitContext{}
	refName := "refs/heads/main"
	reference := plumbing.NewBranchReferenceName(refName)

	ctx = ctx.WithReferenceName(refName)

	assert.Equal(t, reference, ctx.referenceName)
}

func TestGitContext_WithRemoteName(t *testing.T) {
	ctx := &GitContext{}
	remoteName := "origin"

	ctx = ctx.WithRemoteName(remoteName)

	assert.Equal(t, remoteName, ctx.remoteName)
}

func TestGitContext_WithSingleBranch(t *testing.T) {
	ctx := &GitContext{}
	singleBranch := true

	ctx = ctx.WithSingleBranch(singleBranch)

	assert.Equal(t, singleBranch, ctx.singleBranch)
}

func TestGitContext_WithInsecureSkipTLS(t *testing.T) {
	ctx := &GitContext{}
	insecureSkipTLS := true

	ctx = ctx.WithInsecureSkipTLS(insecureSkipTLS)

	assert.Equal(t, insecureSkipTLS, ctx.insecureSkipTLS)
}

func TestGitContext_WithCABundle(t *testing.T) {
	ctx := &GitContext{}
	caBundle := []byte("ca-bundle")

	ctx = ctx.WithCABundle(caBundle)

	assert.Equal(t, caBundle, ctx.caBundle)
}

func TestGitContext_WithTempDir(t *testing.T) {
	ctx := &GitContext{}
	tempDir := "/tmp/test-repo"

	ctx = ctx.WithTempDir(tempDir)

	assert.Equal(t, tempDir, ctx.repositoryPath)
}

func TestGitContext_ExecuteCommand(t *testing.T) {
	ctx, err := InitNewRepository("test-repo")
	defer ctx.Cleanup()
	assert.NoError(t, err)

	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	writeCmd := &WriteFileCommand{FilePath: filePath, Content: content}
	err = ctx.ExecuteCommand(writeCmd)
	assert.NoError(t, err)

	readCmd := &ReadFileCommand{FilePath: filePath}
	err = ctx.ExecuteCommand(readCmd)
	assert.NoError(t, err)
	assert.Equal(t, content, readCmd.Content)
}

func TestGitContext_Cleanup(t *testing.T) {
	ctx, err := InitNewRepository("test-repo")
	assert.NoError(t, err)

	err = ctx.Cleanup()
	assert.NoError(t, err)

	_, err = os.Stat(ctx.repositoryPath)
	assert.True(t, os.IsNotExist(err))
}

func TestGitContext_GetHash(t *testing.T) {
	ctx, err := InitNewRepository("test-repo")
	defer ctx.Cleanup()
	assert.NoError(t, err)

	// Initial commit
	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	writeCmd := &WriteFileCommand{FilePath: filePath, Content: content}
	addCmd := &AddFilesCommand{Files: []string{filePath}}
	commitCmd := &CommitChangesCommand{Message: "Initial commit"}
	err = ctx.ExecuteCommandChain([]GitCommand{writeCmd, addCmd, commitCmd})
	assert.NoError(t, err)

	// Get initial hash
	initialHash, err := ctx.GetHash()
	assert.NoError(t, err)
	assert.NotEmpty(t, initialHash)

	// Create a new commit
	content = []byte("Hello, Git!")
	writeCmd = &WriteFileCommand{FilePath: filePath, Content: content}
	addCmd = &AddFilesCommand{Files: []string{filePath}}
	commitCmd = &CommitChangesCommand{Message: "Second commit"}
	err = ctx.ExecuteCommandChain([]GitCommand{writeCmd, addCmd, commitCmd})
	assert.NoError(t, err)

	// Get new hash
	newHash, err := ctx.GetHash()
	assert.NoError(t, err)
	assert.NotEqual(t, initialHash, newHash)
}

func TestGitContext_GetReferenceName(t *testing.T) {
	ctx, err := InitNewRepository("test-repo")
	defer ctx.Cleanup()
	assert.NoError(t, err)

	// Initial commit
	filePath := "testfile.txt"
	content := []byte("Hello, World!")
	writeCmd := &WriteFileCommand{FilePath: filePath, Content: content}
	addCmd := &AddFilesCommand{Files: []string{filePath}}
	commitCmd := &CommitChangesCommand{Message: "Initial commit"}
	err = ctx.ExecuteCommandChain([]GitCommand{writeCmd, addCmd, commitCmd})
	assert.NoError(t, err)

	// Initial reference name
	initialRefName, err := ctx.GetReferenceName()
	assert.NoError(t, err)
	assert.Equal(t, "refs/heads/master", initialRefName)

	// Create a new branch
	newBranchName := "new-branch"
	err = ctx.repo.CreateBranch(&config.Branch{
		Name:   newBranchName,
		Remote: "origin",
		Merge:  plumbing.NewBranchReferenceName(newBranchName),
	})
	assert.NoError(t, err)

	// Checkout the new branch
	worktree, err := ctx.repo.Worktree()
	assert.NoError(t, err)
	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(newBranchName),
		Create: true,
	})
	assert.NoError(t, err)

	// Get new reference name
	newRefName, err := ctx.GetReferenceName()
	assert.NoError(t, err)
	assert.Equal(t, "refs/heads/"+newBranchName, newRefName)
}
