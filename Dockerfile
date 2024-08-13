FROM golang:1.22.6

RUN go install github.com/cespare/reflex@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["reflex", "-r", "\\.go$", "-s", "go", "run", "main.go"]