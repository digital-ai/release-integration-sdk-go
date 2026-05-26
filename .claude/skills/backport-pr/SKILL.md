---
name: backport-pr
description: Backport PR workflow for release-integration-sdk-go. Use when the user asks to backport a PR, cherry-pick a change, or create a maintenance branch PR.
---

# Backport PR

Use this skill to prepare a backport branch and PR.

## Rules

- Inspect source PR/commit and target branch before changing branches.
- Use non-interactive git commands.
- Run focused tests after cherry-pick conflict resolution.
- Do not force-push or amend unless explicitly requested.

## Command

```bash
bash .claude/skills/backport-pr/router.sh backport <target-branch> <commit-or-pr-ref>
```
