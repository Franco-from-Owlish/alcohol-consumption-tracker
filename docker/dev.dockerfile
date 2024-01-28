FROM golang:1.21-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app
COPY backend .
RUN go mod download

COPY .air.toml .
CMD ["air", "-c", ".air.toml"]