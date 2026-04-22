FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

# Install gcc for CGO/sqlite3
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with CGO
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=1 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata sqlite

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
