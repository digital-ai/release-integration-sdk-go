---
name: release-integration-sdk-go-context
description: release-integration-sdk-go SDK context, integration repositories, Release/Runner relationships, package map, and env-var configurable sibling repo paths. Use when working on SDK changes, plugin compatibility, Release runner behavior, or cross-repo integration checks.
---

# Release Integration SDK Go Context

## Purpose

Use this skill when a task involves this SDK, a plugin that consumes it, Release runner execution, or cross-repo compatibility.

This repository is the Go SDK for Digital.ai Release integration/plugin development. It is plugin-facing API surface, so exported structs, JSON field names, command deserialization, and result behavior must stay compatible unless a breaking change is explicitly requested.

## Architecture Map

```text
Release product
  |
  | dispatches container task payloads
  v
release-runner
  |
  | invokes plugin container / task entrypoint
  v
plugin repo
  |
  | imports github.com/digital-ai/release-integration-sdk-go
  v
SDK packages
  +-- task              payload model, result model, logging helpers
  +-- task/command      command deserialization and execution contract
  +-- runner            simple/command runners and abort handling
  +-- api/release       Release API client wrapper
  +-- api/release/openapi generated Release OpenAPI client/models
  +-- git               Git workspace operations used by integrations
  +-- k8s               Kubernetes client helper
  +-- http              HTTP client helpers
  +-- logger            secret filtering and logger setup
  +-- test              plugin test helpers
```

## Related Repositories

| Env Var | Default | Repository | Purpose |
|---------|---------|------------|---------|
| `RELEASE_K8S_INTEGRATION_PATH` | `../release-k8s-integration` | `digital-ai/release-k8s-integration` | Kubernetes plugin compatibility checks |
| `RELEASE_DEPLOY_INTEGRATION_PATH` | `../release-deploy-integration` | `digital-ai/release-deploy-integration` | Deploy plugin compatibility checks |
| `RELEASE_ARGOCD_INTEGRATION_PATH` | `../release-argocd-integration` | `digital-ai/release-argocd-integration` | Argo CD plugin compatibility checks |
| `RELEASE_FLUXCD_INTEGRATION_PATH` | `../release-fluxcd-integration` | `digital-ai/release-fluxcd-integration` | Flux CD plugin compatibility checks |
| `XL_RELEASE_PATH` | `../../xl-release` | `xebialabs/xl-release` | Release product behavior and payload source |
| `RELEASE_RUNNER_PATH` | `../../release-runner` | `digital-ai/release-runner` | Runner execution and task lifecycle |

Use defaults only as local conveniences. Do not commit machine-specific absolute paths.

## Package Playbooks

| Area | Start Here | Verify With |
|------|------------|-------------|
| Task JSON shape | `task/types.go`, `task/deser.go`, `task/property/deser.go` | `go test ./task/...` |
| Command task regression | `task/command/*.go`, `runner/test/testdata/` | `go test ./task/command/... ./runner/...` |
| Abort/cancel behavior | `runner/executor.go` | `go test ./runner/...` |
| Release API auth/base URL | `api/release/release_client.go` | `go test ./api/release/...` |
| OpenAPI generated model issue | `api/release/openapi/*.go` | Prefer regeneration from source spec |
| Secret filtering | `logger/secret_filter.go` | `go test ./logger/...` |
| Git operations | `git/context.go`, `git/commands.go` | `go test ./git/...` |
| Plugin test helpers | `test/*.go`, `runner/test/*.go` | Check consuming plugin tests |

## Cross-Repo Compatibility Checks

When changing exported SDK APIs or behavior used by plugins:

```bash
RELEASE_K8S_INTEGRATION_PATH=${RELEASE_K8S_INTEGRATION_PATH:-../release-k8s-integration}
RELEASE_DEPLOY_INTEGRATION_PATH=${RELEASE_DEPLOY_INTEGRATION_PATH:-../release-deploy-integration}
RELEASE_ARGOCD_INTEGRATION_PATH=${RELEASE_ARGOCD_INTEGRATION_PATH:-../release-argocd-integration}
RELEASE_FLUXCD_INTEGRATION_PATH=${RELEASE_FLUXCD_INTEGRATION_PATH:-../release-fluxcd-integration}
RELEASE_RUNNER_PATH=${RELEASE_RUNNER_PATH:-../../release-runner}
XL_RELEASE_PATH=${XL_RELEASE_PATH:-../../xl-release}
```

Check plugin usage with content search before changing public names:

```bash
rg "NewCommandRunner|InputContext|ReleaseContext|task.Result|DeserializeCommand" "$RELEASE_K8S_INTEGRATION_PATH" "$RELEASE_DEPLOY_INTEGRATION_PATH" "$RELEASE_ARGOCD_INTEGRATION_PATH" "$RELEASE_FLUXCD_INTEGRATION_PATH"
```

## Rules Of Thumb

- Preserve exported names and JSON tags unless the task explicitly requires a breaking change.
- Add or update fixture tests for deserialization changes.
- Do not edit `api/release/openapi` manually unless the task is specifically about generated API output.
- For runner changes, reason about signal handling, goroutine completion, result channels, and cancellation.
- For logging changes, assume any task input may contain secrets.
- For dependency updates, run `go mod download`, `go build ./...`, and `go test -cover -v ./...`.
