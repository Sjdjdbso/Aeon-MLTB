if [ -f "./server" ]; then
    echo "Starting compiled Go server binary..."
    ./server &
elif command -v go >/dev/null 2>&1 && [ -f "cmd/server/main.go" ]; then
    echo "Building and starting Go server..."
    go build -o server ./cmd/server && ./server &
fi

source .venv/bin/activate
python3 update.py
python3 -m bot
