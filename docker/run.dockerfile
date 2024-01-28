FROM golang:1.21-alpine AS build

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/src
COPY backend .
RUN go mod download

RUN go build -o bin/main cmd/main.go

FROM golang:1.21-alpine AS final

WORKDIR /app

COPY --from=build /app/src/bin/main main

RUN chmod +x main
EXPOSE 8080

CMD ["./main"]