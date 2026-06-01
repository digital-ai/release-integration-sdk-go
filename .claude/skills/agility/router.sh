#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
UPSTREAM_PATH="${AGILITY_AGENT_SKILLS_PATH:-../agility-agent-skills}"

sanitize() {
    printf '%s' "$1" | tr '[:upper:]' '[:lower:]' | tr -cs 'a-z0-9._-' '-' | sed 's/^-//;s/-$//'
}

cmd_fetch() {
    if [ "$#" -lt 1 ]; then
        echo "Usage: $0 fetch <agility-id>"
        return 2
    fi
    if [ -x "$UPSTREAM_PATH/router.sh" ]; then
        "$UPSTREAM_PATH/router.sh" fetch "$@"
        return
    fi
    if [ -x "$UPSTREAM_PATH/scripts/fetch.sh" ]; then
        "$UPSTREAM_PATH/scripts/fetch.sh" "$@"
        return
    fi
    echo "Upstream agility skills not found at $UPSTREAM_PATH"
    echo "Requested item: $1"
}

cmd_branch() {
    if [ "$#" -lt 1 ]; then
        echo "Usage: $0 branch <agility-id> [description]"
        return 2
    fi
    local id="$1"
    shift || true
    local suffix="${*:-work}"
    printf '%s-%s\n' "$(sanitize "$id")" "$(sanitize "$suffix")"
}

cmd_help() {
    cat <<'EOF'
Usage: .claude/skills/agility/router.sh <command> [args]

Commands:
  fetch <id>              Fetch or display an Agility work item
  branch <id> [summary]   Print a sanitized branch name
  help                    Show this help
EOF
}

case "${1:-help}" in
    fetch) shift; cmd_fetch "$@" ;;
    branch) shift; cmd_branch "$@" ;;
    help|*) cmd_help ;;
esac
