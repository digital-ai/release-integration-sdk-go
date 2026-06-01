#!/usr/bin/env bash
set -euo pipefail

VAULT_PATH="${TEAM_MARIO_VAULT_PATH:-../team-mario-agent-vault}"
PRODUCT="${TEAM_MARIO_VAULT_PRODUCT:-plugins}"

sanitize() {
    printf '%s' "$1" | tr '[:upper:]' '[:lower:]' | tr -cs 'a-z0-9._-' '-' | sed 's/^-//;s/-$//'
}

task_dir() {
    if [ "$#" -lt 1 ]; then
        echo "Usage: $0 <command> <task-id> [title]" >&2
        return 2
    fi
    local id title slug
    id="$(sanitize "$1")"
    shift || true
    title="${*:-release-integration-sdk-go-task}"
    slug="$(sanitize "$title")"
    printf '%s/%s/%s-%s\n' "$VAULT_PATH" "$PRODUCT" "$id" "$slug"
}

cmd_init() {
    local dir
    dir="$(task_dir "$@")"
    mkdir -p "$dir"
    if [ ! -f "$dir/spec.md" ]; then
        cat > "$dir/spec.md" <<'EOF'
# Spec

## Problem

## Assumptions

## Success Criteria

## Out of Scope
EOF
    fi
    if [ ! -f "$dir/implementation-plan.md" ]; then
        cat > "$dir/implementation-plan.md" <<'EOF'
# Implementation Plan

## Approach

## Files To Inspect

## Planned Changes

## Verification
EOF
    fi
    echo "$dir"
}

cmd_notes() {
    local dir
    dir="$(task_dir "$@")"
    mkdir -p "$dir"
    if [ ! -f "$dir/session-notes.md" ]; then
        cat > "$dir/session-notes.md" <<'EOF'
# Session Notes

## Summary

## What Changed

## Verification

## PR / Links

## Follow-Ups
EOF
    fi
    echo "$dir/session-notes.md"
}

cmd_help() {
    cat <<'EOF'
Usage: .claude/skills/team-mario-vault/router.sh <command> <task-id> [title]

Commands:
  init <task-id> [title]    Create spec.md and implementation-plan.md
  notes <task-id> [title]   Create session-notes.md
  path <task-id> [title]    Print the vault task directory
  help                      Show this help
EOF
}

case "${1:-help}" in
    init) shift; cmd_init "$@" ;;
    notes) shift; cmd_notes "$@" ;;
    path) shift; task_dir "$@" ;;
    help|*) cmd_help ;;
esac
