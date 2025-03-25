package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"path/filepath"
	"sync"
)

// GitCommand is an interface that defines a method to execute a command on a git repository.
type GitCommand interface {
	execute(repo *GitContext) error
}

// GitContext is a struct that contains the context for a git repository.
type GitContext struct {
	repo            *git.Repository
	repositoryPath  string
	authMethod      transport.AuthMethod
	repoURL         string
	proxyURL        string
	proxyUsername   string
	proxyPassword   string
	remoteName      string
	referenceName   plumbing.ReferenceName
	singleBranch    bool
	insecureSkipTLS bool
	caBundle        []byte
	once            sync.Once
}

// NewRepository creates a new GitContext with the given repository URL and authentication method.
func NewRepository(repoURL string) *GitContext {
	repositoryName := filepath.Base(repoURL)
	return &GitContext{
		repositoryPath: os.TempDir() + "/" + repositoryName,
		repoURL:        repoURL,
	}
}

// InitNewRepository initializes a new local GitContext with the given repository name.
func InitNewRepository(repoName string) (*GitContext, error) {
	var err error
	repoPath := filepath.Join(os.TempDir(), repoName)
	ctx := &GitContext{
		repositoryPath: repoPath,
	}
	ctx.repo, err = git.PlainInit(ctx.repositoryPath, false)
	return ctx, err
}

// WithToken sets the token for the GitContext.
func (ctx *GitContext) WithToken(token string) *GitContext {
	ctx.authMethod = &http.TokenAuth{Token: token}
	return ctx
}

// WithBasicAuth sets the username and password for the GitContext.
func (ctx *GitContext) WithBasicAuth(username, password string) *GitContext {
	ctx.authMethod = &http.BasicAuth{Username: username, Password: password}
	return ctx
}

// WithProxy sets the proxy URL, username, and password for the GitContext.
func (ctx *GitContext) WithProxy(proxyURL, proxyUsername, proxyPassword string) *GitContext {
	ctx.proxyURL = proxyURL
	ctx.proxyUsername = proxyUsername
	ctx.proxyPassword = proxyPassword
	return ctx
}

// WithReferenceName sets the reference name for the GitContext.
func (ctx *GitContext) WithReferenceName(referenceName string) *GitContext {
	ctx.referenceName = plumbing.NewBranchReferenceName(referenceName)
	return ctx
}

// WithRemoteName sets the remote name for the GitContext. Default is "origin".
func (ctx *GitContext) WithRemoteName(remoteName string) *GitContext {
	ctx.remoteName = remoteName
	return ctx
}

// WithSingleBranch sets the single branch flag for the GitContext. Default is false. If true, only the WithReferenceName set branch is cloned.
func (ctx *GitContext) WithSingleBranch(singleBranch bool) *GitContext {
	ctx.singleBranch = singleBranch
	return ctx
}

// WithInsecureSkipTLS sets the insecure skip TLS flag for the GitContext. Default is false. If true, check of the certificate is skipped, not recommended in production.
func (ctx *GitContext) WithInsecureSkipTLS(insecureSkipTLS bool) *GitContext {
	ctx.insecureSkipTLS = insecureSkipTLS
	return ctx
}

// WithCABundle sets the CA bytes for the GitContext. CA will be used as the certificate of trust for the TLS connection.
func (ctx *GitContext) WithCABundle(caBundle []byte) *GitContext {
	ctx.caBundle = caBundle
	return ctx
}

// WithTempDir sets the temporary directory for the GitContext. Default is the OS temporary directory.
func (ctx *GitContext) WithTempDir(tempDir string) *GitContext {
	ctx.repositoryPath = tempDir
	return ctx
}

// ExecuteCommand executes the given GitCommand on the GitContext.
func (ctx *GitContext) ExecuteCommand(cmd GitCommand) error {
	var err error
	ctx.once.Do(func() {
		if ctx.repo == nil {

			cloneOptions := &git.CloneOptions{
				URL:             ctx.repoURL,
				Depth:           1,
				RemoteName:      ctx.remoteName,
				ReferenceName:   ctx.referenceName,
				SingleBranch:    ctx.singleBranch,
				InsecureSkipTLS: ctx.insecureSkipTLS,
				CABundle:        ctx.caBundle,
				Auth:            ctx.authMethod,
			}
			if ctx.proxyURL != "" {
				cloneOptions.ProxyOptions = transport.ProxyOptions{
					URL:      ctx.proxyURL,
					Username: ctx.proxyUsername,
					Password: ctx.proxyPassword,
				}
			}

			ctx.repo, err = git.PlainClone(ctx.repositoryPath, false, cloneOptions)
		}
	})
	if err != nil {
		return err
	}

	return cmd.execute(ctx)
}

// ExecuteCommandChain executes the given chain of GitCommands on the GitContext. If any of the commands fail, the chain is stopped and the error is returned. Otherwise, nil is returned. The commands are executed in the order they are given. When command fails the chain is stopped, but previously executed commands are not reverted.
func (ctx *GitContext) ExecuteCommandChain(cmds []GitCommand) error {
	for _, cmd := range cmds {
		err := ctx.ExecuteCommand(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Cleanup removes the temporary directory of the GitContext. This should be called after the GitContext is no longer needed. Use defer ctx.Cleanup() after creating a GitContext.
func (ctx *GitContext) Cleanup() error {
	return os.RemoveAll(ctx.repositoryPath)
}
