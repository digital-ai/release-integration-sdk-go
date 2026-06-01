# AGENTS.md - Go SDK Agent Guide

## Default Workflow

1. Inspect the relevant files before editing.
2. State assumptions when behavior is ambiguous.
3. Make the smallest correct change.
4. Run the narrowest useful verification first, then broader checks when warranted.
5. Do not modify generated Release OpenAPI files unless the task is specifically about generated API output.

## Go Style

| Rule | Guidance |
|------|----------|
| Formatting | Run `gofmt` on edited Go files. Do not hand-format. |
| Imports | Let `gofmt`/`goimports` style guide import grouping; preserve existing grouping if not using `goimports`. |
| Errors | Return errors with context. Avoid swallowing errors. |
| Tests | Prefer table tests for multiple cases. Keep fixtures under existing `testdata/` directories. |
| APIs | Preserve exported API compatibility unless the task explicitly requires a breaking change. |
| Logging | Do not log secrets. Use existing logger and secret filter patterns. |
| Context | Preserve cancellation semantics in runner and command code. |

## Verification Commands

| Change Type | Command |
|-------------|---------|
| Any Go change | `go test ./...` |
| CI parity | `go mod download && go build ./... && go test -cover -v ./...` |
| Formatting | `gofmt -w <edited-go-files>` |
| Static checks | `go vet ./...` |
| Vulnerabilities | `govulncheck ./...` |

## Package Notes

| Package | Notes |
|---------|-------|
| `runner` | Be careful with goroutines, signals, abort handling, and result channel behavior. |
| `task` | This is plugin-facing API surface. Avoid breaking JSON field names or exported structs. |
| `task/command` | Deserialization impacts command-based plugin tasks. Add fixture tests for regressions. |
| `api/release/openapi` | Generated client/models. Prefer regenerating from source spec when broad changes are needed. |
| `git` | Tests may use local repositories and temporary directories. Preserve cross-platform paths. |
| `logger` | Secret filtering must be conservative and covered by tests. |
| `test` | Shared SDK test helpers used by plugin repositories. Keep compatibility in mind. |

## Dependency Changes

- Keep `go.mod` and `go.sum` consistent.
- After dependency updates, run `go mod tidy` only when it is directly relevant.
- For SDK version bumps affecting plugins, check representative plugin repositories using env-var-configurable paths from the context skill.

## PR Expectations

- Explain SDK API impact, plugin compatibility impact, and verification performed.
- Include tests for behavior changes.
- For vulnerability PRs, include dependency before/after and `govulncheck` output summary.
