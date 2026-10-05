#!/usr/bin/env bash
set -e

if [ -f "cmd/server/main.go" ]; then
    echo "Building Go server..."
    go build -o server ./cmd/server/main.go
fi

if [ -f "server" ]; then
    echo "Starting Go server in background..."
    GOMEMLIMIT=100MiB ./server &
fi

if [ -d ".venv" ]; then
    source .venv/bin/activate
fi

python3 update.py
python3 -m bot
