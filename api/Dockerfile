# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO無効で静的リンク（必要に応じて）
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /out/app ./...

# Run stage（distrolessでもOK。ここではalpine最小で）
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /out/app /app/app
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app/app"]
