---
name: create-pr
description: Pull request creation workflow for release-integration-sdk-go. Use when creating PRs, preparing PR descriptions, checking branch diffs, or publishing GitHub pull requests.
---

# Create PR

Use this skill when the user asks to create or prepare a PR.

## Rules

- Inspect `git status`, `git diff`, and recent commits before creating a PR.
- Confirm only intended files are included.
- Include SDK API impact, plugin compatibility impact, and verification performed.
- Use `gh` for GitHub operations.

## Command

```bash
bash .claude/skills/create-pr/router.sh create
```

Pass additional `gh pr create` arguments after `create` when needed.
