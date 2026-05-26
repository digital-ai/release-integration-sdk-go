---
name: agility
description: Agility work item helper and branch naming wrapper. Use when the user mentions Agility, story, defect, S-number, D-number, start branch, or work item metadata.
---

# Agility Wrapper

This skill is a thin wrapper around `digital-ai/agility-agent-skills` when that repository is available locally. Keep this wrapper small so upstream improvements can be consumed without forking behavior.

## Configuration

| Env Var | Default | Purpose |
|---------|---------|---------|
| `AGILITY_AGENT_SKILLS_PATH` | `../agility-agent-skills` | Local checkout of `digital-ai/agility-agent-skills` |

## Commands

```bash
bash .claude/skills/agility/router.sh help
bash .claude/skills/agility/router.sh fetch S-12345
bash .claude/skills/agility/router.sh branch S-12345
```

If the upstream checkout is not available, the wrapper still provides deterministic branch-name output.
