source .venv/bin/activate
if command -v go >/dev/null 2>&1; then
    go build -o server ./cmd/server
    ./server > go_server.log 2>&1 &
fi
python3 update.py
python3 -m bot
