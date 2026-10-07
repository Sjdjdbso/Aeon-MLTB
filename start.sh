source .venv/bin/activate
if [ -f "go.mod" ]; then
    go build -o aeon-server ./cmd/server && ./aeon-server &
fi
python3 update.py
python3 -m bot
