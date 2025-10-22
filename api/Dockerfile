# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO無効で静的リンク（必要に応じて）
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /out/app ./...

# Dev stage（開発用：Debian 12 + Go + air + sh あり）
FROM golang:1.25-alpine AS dev
WORKDIR /app
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPATH=/go
RUN apk add --no-cache git bash build-base

# 依存だけ先に解決してビルドキャッシュを最適化
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

#ENV PORT=8080
#EXPOSE 8080

# air をインストール（PATH: /go/bin）
#RUN go install github.com/air-verse/air@latest
#CMD ["sh", "-lc", "cd /app && [ -f go.mod ] && go mod download; exec air -c .air.toml"]

# デフォルトは air 起動（Composeで上書きしてもOK）
#CMD ["air", "-c", ".air.toml"]


# Run stage（distrolessでもOK。ここではalpine最小で）
FROM gcr.io/distroless/base-debian12 AS runner
WORKDIR /app
COPY --from=builder /out/app /app/app
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app/app"]
