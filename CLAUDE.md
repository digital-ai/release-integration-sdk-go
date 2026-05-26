# CLAUDE.md - Release Integration SDK Go Development Harness

> Read `AGENTS.md` for Go style and workflow rules.
> Load the `release-integration-sdk-go-context` skill for the SDK and integration map.

---

## Principles

1. **Think before coding** - State assumptions. Surface ambiguity. Ask before guessing when behavior is unclear.
2. **Simplicity first** - Write the minimum code that solves the problem. Avoid speculative abstractions.
3. **Surgical changes** - Touch only what you must. Match existing style. Clean up only your own changes.
4. **Goal-driven execution** - Define success criteria, verify each step, and loop until verified.

---

## What This Repo Is

`release-integration-sdk-go` is the Go SDK used by Digital.ai Release integrations and plugins that execute containerized tasks. It provides task input/output types, runner execution helpers, Release API clients, Git helpers, Kubernetes helpers, HTTP utilities, logging, and test fixtures used by plugin repositories.

```text
Release / Runner
      |
      v
Plugin container task
      |
      v
release-integration-sdk-go
      |
      +-- task/       task input, result, security, CI/deploy types
      +-- runner/     runner lifecycle, abort handling, command execution
      +-- api/        Release OpenAPI client wrapper and generated models
      +-- git/        repository read/write/commit helpers
      +-- k8s/        Kubernetes client helper
      +-- http/       HTTP client builders
      +-- logger/     secret filtering and task logging
      +-- test/       test DSL and mock remotes
```

---

## Supported Versions

| Branch | Version |
|--------|---------|
| `master` | current SDK line |
| `release/*` | release branches when present |

Tags are created as `v*` by CI on `master`.

---

## Module Map

| Area | Key Files |
|------|-----------|
| Module definition | `go.mod`, `go.sum` |
| Runner execution | `runner/executor.go`, `runner/watcher.go` |
| Task model | `task/types.go`, `task/result.go`, `task/deser.go`, `task/command/*.go` |
| Release API | `api/release/release_client.go`, `api/release/openapi/*.go` |
| Git helpers | `git/context.go`, `git/commands.go` |
| Kubernetes helpers | `k8s/default-client.go` |
| HTTP helpers | `http/http_client.go`, `http/http_client_builder.go` |
| Logging and secrets | `logger/logger.go`, `logger/secret_filter.go` |
| Test support | `test/*.go`, `runner/test/*.go` |

---

## Related Repositories

| Repository | Purpose |
|------------|---------|
| `digital-ai/release-k8s-integration` | Kubernetes Release plugin using this SDK |
| `digital-ai/release-deploy-integration` | Deploy integration plugin using this SDK |
| `digital-ai/release-argocd-integration` | Argo CD Release plugin using this SDK |
| `digital-ai/release-fluxcd-integration` | Flux CD Release plugin using this SDK |
| `xebialabs/xl-release` | Digital.ai Release product |
| `digital-ai/release-runner` | Release runner that executes integration tasks |

Use env vars for local cross-repo work. Do not hardcode machine-specific paths in committed docs or scripts.

---

## Documentation References

- SDK repository: https://github.com/digital-ai/release-integration-sdk-go
- SDK wiki: https://github.com/digital-ai/release-integration-sdk-go/wiki
- Release docs: https://docs.digital.ai/
- Release docs source: https://github.com/digital-ai/docs-release
- Runner repository: https://github.com/digital-ai/release-runner

---

## Local Dev Quick Start

| Goal | Command |
|------|---------|
| Download dependencies | `go mod download` |
| Build everything | `go build ./...` |
| Run all tests | `go test -cover -v ./...` |
| Run one package | `go test -v ./runner/...` |
| Format Go files | `gofmt -w <files>` |
| Check vulnerabilities | `govulncheck ./...` |

---

## Common Task Playbooks

| Problem | Start Here |
|---------|------------|
| Plugin task input fails to deserialize | `task/deser.go`, `task/property/deser.go`, relevant plugin task JSON |
| Command execution or abort behavior is wrong | `runner/executor.go`, `task/command/*.go`, `runner/test/testdata/` |
| Release API request/auth issue | `api/release/release_client.go`, generated `api/release/openapi/api_*.go` |
| Secret appears in logs | `logger/secret_filter.go`, `logger/secret_filter_test.go` |
| Git-backed task behavior fails | `git/context.go`, `git/commands.go`, `git/*_test.go` |
| Kubernetes client setup fails | `k8s/default-client.go`, consuming plugin repository |
| SDK update breaks a plugin | Load `release-integration-sdk-go-context`; compare plugin usage of changed SDK APIs |

---

## Vault Workflow

For every non-trivial task:

1. Before implementation, write `spec.md` and `implementation-plan.md` in the Team Mario vault.
2. During implementation, keep changes surgical and verification-focused.
3. After PR creation or handoff, write `session-notes.md` with what changed, what worked, verification, and links.

Use the `team-mario-vault` skill and commands for vault operations.

---

## Available Skills

| Skill | Purpose |
|-------|---------|
| `release-integration-sdk-go-context` | SDK architecture and related repo map |
| `agility` | Wrapper around Digital.ai agility agent skills |
| `team-mario-vault` | Vault task docs and session notes |
| `vulnerability-report` | Go module vulnerability/dependency report |
| `create-pr` | PR creation workflow |
| `backport-pr` | Backport workflow |

## Available Commands

| Command | Purpose |
|---------|---------|
| `/build` | Run `go mod download` and `go build ./...` |
| `/test` | Run Go tests, optionally with args |
| `/lint` | Run `go vet ./...` |
| `/fmt` | Format Go files |
| `/run-dev` | Run a package or command passed as arguments |
| `/create-pr` | Create a PR via the `create-pr` skill |
| `/backport-pr` | Backport a PR via the `backport-pr` skill |
| `/agility-fetch` | Fetch an Agility item via the `agility` skill |
| `/agility-branch` | Create a work branch from an Agility item |
| `/vault-init` | Create vault spec and implementation plan |
| `/vault-notes` | Create vault session notes |
| `/vulnerability-report` | Generate vulnerability/dependency report |
