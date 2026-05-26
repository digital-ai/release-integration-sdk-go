#!/usr/bin/env bash
set -euo pipefail

sanitize() {
    printf '%s' "$1" | tr '[:upper:]' '[:lower:]' | tr -cs 'a-z0-9._-' '-' | sed 's/^-//;s/-$//'
}

cmd_backport() {
    if [ "$#" -lt 2 ]; then
        echo "Usage: $0 backport <target-branch> <commit-sha>" >&2
        return 2
    fi
    local target="$1"
    local ref="$2"
    local branch="backport-$(sanitize "$ref")-to-$(sanitize "$target")"
    git fetch origin "$target"
    git switch -c "$branch" "origin/$target"
    git cherry-pick "$ref"
    echo "Created backport branch: $branch"
}

cmd_help() {
    cat <<'EOF'
Usage: .claude/skills/backport-pr/router.sh <command> [args]

Commands:
  backport <target-branch> <commit-sha>  Create branch and cherry-pick commit
  help                                  Show this help
EOF
}

case "${1:-help}" in
    backport) shift; cmd_backport "$@" ;;
    help|*) cmd_help ;;
esac
