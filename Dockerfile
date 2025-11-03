# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags '-extldflags "-static" -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}' \
    -o tabloom ./cmd/tabloom

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

RUN adduser -D -s /bin/sh tabloom
USER tabloom

COPY --from=builder --chown=tabloom:tabloom /app/tabloom .
COPY --from=builder --chown=tabloom:tabloom /app/migrations ./migrations

EXPOSE 8080

CMD ["./tabloom"]