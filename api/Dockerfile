# FROM golang:1.25-alpine AS base
# WORKDIR /app
# ENV CGO_ENABLED=0 \
#     GO111MODULE=on \
#     GOPATH=/go
# RUN apk add --no-cache git bash build-base

# # 依存だけ先に解決してビルドキャッシュを最適化
# COPY go.mod go.sum ./
# RUN --mount=type=cache,target=/go/pkg/mod \
#     go mod download

# # ===== dev =====
# FROM base AS dev
# # air を導入（cosmtrek/air）
# RUN go install github.com/air-verse/air@v1.52.3
# # ソース一式（ホットリロード時はボリュームで上書きされる）
# COPY . .
# # デフォルトは air 起動
# CMD ["air", "-c", ".air.toml"]


# ENV PORT=8080
# EXPOSE 8080


# # Run stage（distrolessでもOK。ここではalpine最小で）
# FROM gcr.io/distroless/base-debian12 AS runner
# WORKDIR /app
# COPY --from=builder /out/app /app/app
# ENV PORT=8080
# EXPOSE 8080
# USER nonroot:nonroot
# CMD ["/app/app"]



# ===== base (Debian bookworm, Go 1.25) =====
FROM golang:1.25-bookworm AS base
WORKDIR /app
ENV CGO_ENABLED=0 \
    #GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct \
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