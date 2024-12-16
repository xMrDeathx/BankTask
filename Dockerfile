# syntax=docker/dockerfile:1

# Stage 1: Build
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum config.env . ./
RUN go mod download
RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
RUN go get github.com/joho/godotenv
RUN go install github.com/joho/godotenv/cmd/godotenv@latest
RUN go get github.com/jackc/pgx/v4
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY . ./
RUN oapi-codegen -generate gorilla,types wallet/api/frontend/frontendapi.yaml > wallet/api/frontend/frontendapi.gen.go
RUN CGO_ENABLED=0 go build -o app -buildvcs=false main.go

# Stage 2: Runtime
FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /app/config.env ./
COPY --from=builder /app/wallet/db/migrations ./wallet/db/migrations
CMD ["./app"]