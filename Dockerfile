FROM golang:1.26.4
WORKDIR /app
COPY . .

CMD ["go", "run", "main.go"]
