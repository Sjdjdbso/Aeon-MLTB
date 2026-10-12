source .venv/bin/activate
if [ -f "cmd/server/main.go" ]; then
    go build -o server ./cmd/server
fi
if [ -f "server" ]; then
    ./server &
fi
python3 update.py
python3 -m bot
