source .venv/bin/activate

if command -v go >/dev/null 2>&1; then
    echo "Building Go web server..."
    go build -o bin/webserver ./cmd/server
    if [ -f bin/webserver ]; then
        echo "Starting Go web server in background..."
        ./bin/webserver &
    fi
elif [ -f bin/webserver ]; then
    echo "Starting precompiled Go web server in background..."
    ./bin/webserver &
fi

python3 update.py
python3 -m bot