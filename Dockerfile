FROM 5hojib/aeon:latest

WORKDIR /app
RUN chmod 777 /app

RUN uv venv
COPY requirements.txt .
RUN uv pip install --no-cache-dir -r requirements.txt

COPY . .
RUN apt-get update && apt-get install -y golang-go && rm -rf /var/lib/apt/lists/*
RUN go build -o aeon-server ./cmd/server
CMD ["bash", "start.sh"]
