package git

import (
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
	ctx := NewRepository(repoURL, auth)

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
	refName := plumbing.NewBranchReferenceName("refs/heads/main")

	ctx = ctx.WithReferenceName(refName)

	assert.Equal(t, refName, ctx.referenceName)
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
