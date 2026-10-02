FROM 5hojib/aeon:latest

WORKDIR /app
RUN chmod 777 /app

RUN apt-get update && apt-get install -y golang-go && rm -rf /var/lib/apt/lists/*

RUN uv venv
COPY requirements.txt .
RUN uv pip install --no-cache-dir -r requirements.txt

COPY . .
RUN go build -o bin/webserver ./cmd/server

CMD ["bash", "start.sh"]