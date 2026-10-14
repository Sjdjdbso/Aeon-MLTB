source .venv/bin/activate
if [ -f "./aeon-server" ]; then
    ./aeon-server > aeon-server.log 2>&1 &
fi
python3 update.py
python3 -m bot