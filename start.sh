#!/bin/bash
if [ -d ".venv" ]; then
    source .venv/bin/activate
fi

if [ -f "cmd/server/main.go" ]; then
    echo "Building and starting Go server..."
    go build -o aeon-server ./cmd/server && ./aeon-server &
fi

python3 update.py
python3 -m bot
