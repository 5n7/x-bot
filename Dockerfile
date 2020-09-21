FROM golang:1.15-alpine AS builder

RUN apk --update add make

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build
RUN go get github.com/githubnemo/CompileDaemon

FROM alpine

COPY --from=builder /app/bin/x-bot /usr/bin/x-bot
CMD ["x-bot"]
