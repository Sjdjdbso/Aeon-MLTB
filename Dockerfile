FROM 5hojib/aeon:latest

WORKDIR /app
RUN chmod 777 /app

RUN uv venv
COPY requirements.txt .
RUN uv pip install --no-cache-dir -r requirements.txt

COPY . .
RUN go build -o go_server ./cmd/server
CMD ["bash", "start.sh"]
