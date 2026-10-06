#!/bin/bash
if [ -f "cmd/server/main.go" ]; then
    echo "Building Go server..."
    mkdir -p bin
    go build -o bin/server ./cmd/server
    if [ -f "bin/server" ]; then
        echo "Starting Go server in background..."
        bin/server > go_server.log 2>&1 &
    fi
fi

source .venv/bin/activate
python3 update.py
python3 -m bot
