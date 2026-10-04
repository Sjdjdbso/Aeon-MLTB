if [ -f "cmd/server/main.go" ]; then
    go build -o aeon-mltb ./cmd/server/main.go && ./aeon-mltb &
elif [ -f "./aeon-mltb" ]; then
    ./aeon-mltb &
fi

source .venv/bin/activate
python3 update.py
python3 -m bot
