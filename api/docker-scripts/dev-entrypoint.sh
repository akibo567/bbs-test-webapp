#!/usr/bin/env sh
set -eu

cd /app

# 1) go.mod が無ければ初期化
#if [ ! -f go.mod ]; then
#  MODPATH="${MODULE_PATH:-example.com/helloserver}"
#  echo "🧩 go mod init ${MODPATH}"
#  go mod init "${MODPATH}"
#fi

# 2) go の互換バージョンを固定したい場合（例: 1.25）
#if [ "${GO_COMPAT:-}" != "" ]; then
#  echo "🔧 go mod edit -go=${GO_COMPAT}"
#  go mod edit -go="${GO_COMPAT}"
#fi

# 3) 依存整理（初回で go.sum が無くても2回目で救済）
echo "📦 go mod tidy"
if ! go mod tidy; then
  echo "⚠️ go mod tidy failed once. Touching go.sum and retrying..."
  touch go.sum
  go mod tidy || true
fi

# 4) 本命（Air など）を実行
exec "$@"