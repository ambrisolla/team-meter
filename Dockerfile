FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app cmd/jira-issues/main.go

FROM debian:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app/

COPY --from=builder /app/app .
COPY --from=builder /app/config/jira_projects.txt ./config/
RUN chmod +x /app/app

CMD ["/app/app"]
