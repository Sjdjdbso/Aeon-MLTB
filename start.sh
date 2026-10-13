#!/bin/bash
source .venv/bin/activate

export GO_PORT="${GO_PORT:-8075}"

if [ ! -f "./aeon-server" ] && command -v go &> /dev/null; then
    echo "Building Go microservice..."
    go build -o aeon-server ./cmd/server
fi

if [ -f "./aeon-server" ]; then
    echo "Starting Go microservice on port ${GO_PORT} in background..."
    ./aeon-server &
fi

python3 update.py
python3 -m bot
