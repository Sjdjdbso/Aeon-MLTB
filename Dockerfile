FROM golang:1.24-alpine AS go-builder
WORKDIR /app
COPY go.mod ./
COPY cmd/ ./cmd/
RUN go build -o server ./cmd/server

FROM 5hojib/aeon:latest

WORKDIR /app
RUN chmod 777 /app

RUN uv venv
COPY requirements.txt .
RUN uv pip install --no-cache-dir -r requirements.txt

COPY . .
COPY --from=go-builder /app/server ./server

CMD ["bash", "start.sh"]
