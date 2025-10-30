ARG GOLANG_IMAGE=golang:1.25.3-bookworm


# ===== base =====
FROM ${GOLANG_IMAGE} AS base
WORKDIR /app
ENV CGO_ENABLED=0 \
    #GO111MODULE=on \
    GOPROXY=https://proxy.golang.org \
    GOSUMDB=sum.golang.org \
    GOTOOLCHAIN=local

# 開発ユーティリティと証明書
RUN apt-get update && apt-get install -y --no-install-recommends \
      git ca-certificates bash build-essential \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates

# ===== dev (hot reload with Air) =====
FROM base AS dev
RUN go install github.com/air-verse/air@latest

#COPY go.mod go.sum ./
#RUN go mod download
# プロジェクト一式
#COPY . .
# 起動時に go mod init/tidy を自動化するエントリポイント
COPY docker-scripts/dev-entrypoint.sh /usr/local/bin/dev-entrypoint
RUN chmod +x /usr/local/bin/dev-entrypoint
ENTRYPOINT ["/usr/local/bin/dev-entrypoint"]
CMD ["air", "-c", ".air.toml"]

ENV PORT=8080
EXPOSE 8080

# ===== builder (prod) =====
FROM base AS builder
COPY . .
# ビルドキャッシュ活用
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /bin/app ./

# ===== prod (distroless / glibc系 / 最小) =====
FROM gcr.io/distroless/static-debian12:nonroot AS prod
WORKDIR /app
COPY --from=builder /bin/app /app/app
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]