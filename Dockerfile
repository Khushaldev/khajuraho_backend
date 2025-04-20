FROM golang:1.24-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache bash
RUN go install github.com/air-verse/air@latest
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN go mod tidy

# RUN go build -gcflags="all=-N -l" -o server ./cmd/server

# EXPOSE 3000
# EXPOSE 2345

# CMD ["air", "cmd/server/main.go", "-b", "0.0.0.0"]
# CMD ["dlv", "exec", "cmd/server/main.go", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient"]


