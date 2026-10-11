FROM 5hojib/aeon:latest

WORKDIR /app
RUN chmod 777 /app

RUN uv venv
COPY requirements.txt .
RUN uv pip install --no-cache-dir -r requirements.txt

COPY . .
RUN if command -v go >/dev/null 2>&1 && [ -f "cmd/server/main.go" ]; then go build -o server ./cmd/server; fi

CMD ["bash", "start.sh"]
