#!/usr/bin/env bash
set -euo pipefail

cmd_status() {
    git status --short
    git log --oneline -10
}

cmd_create() {
    git status --short
    git log --oneline -10
    if ! command -v gh >/dev/null 2>&1; then
        echo "gh is not installed or not on PATH" >&2
        return 127
    fi
    if [ "$#" -gt 0 ]; then
        gh pr create "$@"
    else
        gh pr create --fill
    fi
}

cmd_help() {
    cat <<'EOF'
Usage: .claude/skills/create-pr/router.sh <command> [gh-pr-create-args]

Commands:
  status    Show current status and recent commits
  create    Run gh pr create, defaulting to --fill
  help      Show this help
EOF
}

case "${1:-help}" in
    status) shift; cmd_status "$@" ;;
    create) shift; cmd_create "$@" ;;
    help|*) cmd_help ;;
esac
