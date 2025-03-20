package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"os"
	"path/filepath"
	"sync"
)

type GitCommand interface {
	execute(repo *git.Repository) error
}

type AuthMethod interface {
	transport.AuthMethod
}

type GitContext struct {
	repo            *git.Repository
	repositoryPath  string
	authMethod      AuthMethod
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

func NewRepository(repoURL string, auth AuthMethod) *GitContext {
	repositoryName := filepath.Base(repoURL)
	return &GitContext{
		repositoryPath: os.TempDir() + "/" + repositoryName,
		repoURL:        repoURL,
		authMethod:     auth,
	}
}

func InitNewRepository(repoName string) (*GitContext, error) {
	var err error
	repoPath := filepath.Join(os.TempDir(), repoName)
	ctx := &GitContext{
		repositoryPath: repoPath,
	}
	ctx.repo, err = git.PlainInit(ctx.repositoryPath, false)
	return ctx, err
}

func (ctx *GitContext) WithProxy(proxyURL, proxyUsername, proxyPassword string) *GitContext {
	ctx.proxyURL = proxyURL
	ctx.proxyUsername = proxyUsername
	ctx.proxyPassword = proxyPassword
	return ctx
}

func (ctx *GitContext) WithReferenceName(referenceName plumbing.ReferenceName) *GitContext {
	ctx.referenceName = referenceName
	return ctx
}

func (ctx *GitContext) WithRemoteName(remoteName string) *GitContext {
	ctx.remoteName = remoteName
	return ctx
}

func (ctx *GitContext) WithSingleBranch(singleBranch bool) *GitContext {
	ctx.singleBranch = singleBranch
	return ctx
}

func (ctx *GitContext) WithInsecureSkipTLS(insecureSkipTLS bool) *GitContext {
	ctx.insecureSkipTLS = insecureSkipTLS
	return ctx
}

func (ctx *GitContext) WithCABundle(caBundle []byte) *GitContext {
	ctx.caBundle = caBundle
	return ctx
}

func (ctx *GitContext) WithTempDir(tempDir string) *GitContext {
	ctx.repositoryPath = tempDir
	return ctx
}

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

	return cmd.execute(ctx.repo)
}

func (ctx *GitContext) Cleanup() error {
	return os.RemoveAll(ctx.repositoryPath)
}
