source .venv/bin/activate
if [ -f "./go_server" ]; then
    ./go_server &
elif [ -f "./cmd/server/main.go" ]; then
    go run ./cmd/server &
fi
python3 update.py
python3 -m bot
