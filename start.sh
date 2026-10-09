source .venv/bin/activate

if command -v go >/dev/null 2>&1; then
    echo "Building Go server..."
    go build -o server ./cmd/server
    if [ -f "./server" ]; then
        echo "Starting Go server in background..."
        ./server > go_server.log 2>&1 &
    fi
elif [ -f "./server" ]; then
    echo "Starting compiled Go server in background..."
    ./server > go_server.log 2>&1 &
fi

python3 update.py
python3 -m bot
