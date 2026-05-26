---
name: team-mario-vault
description: Team Mario vault workflow for spec.md, implementation-plan.md, and session-notes.md. Use for non-trivial tasks, handoffs, PR notes, and work item documentation.
---

# Team Mario Vault

Use this skill to create and update task documentation in the Team Mario agent vault.

## Configuration

| Env Var | Default | Purpose |
|---------|---------|---------|
| `TEAM_MARIO_VAULT_PATH` | `../team-mario-agent-vault` | Local vault checkout |
| `TEAM_MARIO_VAULT_PRODUCT` | `plugins` | Product folder for SDK/plugin work |

## Workflow

Before implementation:

```bash
bash .claude/skills/team-mario-vault/router.sh init S-12345 short-title
```

After PR or handoff:

```bash
bash .claude/skills/team-mario-vault/router.sh notes S-12345 short-title
```

Generated files are templates. Fill in concrete assumptions, plan, verification, and links.
