#!/usr/bin/env bash
# Consume a fresh public rendered fixture; never reimplement generation.
set -euo pipefail
project="$1"
cd "$project"
for config in .env.example .env.docker; do
  if rg -n 'APP_HOST=|APP_PORT=|TEMPORAL__|cmd/app' "$config"; then
    echo "stale environment contract" >&2
    exit 1
  fi
done
go test ./...
go build ./...
check_dir="$(mktemp -d)"
trap 'rm -rf "$check_dir"' EXIT
go build -o "$check_dir/service" ./cmd/service
if APP_LISTEN_ADDR='invalid::address' "$check_dir/service" >"$check_dir/startup.log" 2>&1; then
  echo "invalid listen address returned success" >&2
  exit 1
fi
grep -q 'http server' "$check_dir/startup.log"
echo "render checks passed: tests, build, invalid HTTP startup exits nonzero"
