# AGENTS.md

このリポジトリで作業するAIエージェント向けの最小ガイドです。

## 概要
- Next.js(Frontend) + Go(Backend) + PostgreSQL + nginx を Docker Compose で動かす雛形
- 開発/確認環境はシェルスクリプト経由で起動する前提

## よく使うコマンド
開発環境:
- 初回ビルド: `./dev_build.sh`
- 起動: `./dev_run.sh`
- 停止: `./dev_stop.sh`

確認環境:
- 初回ビルド: `./prod_build.sh`
- 起動: `./product_run.sh`
- 停止: `./product_stop.sh`

## 主要ディレクトリ
- `frontend/`: Next.js フロントエンド
- `api/`: Go バックエンド
- `db/`: DB関連(初期化スクリプト等)
- `nginx/`: nginx 設定
- `docker-compose*.yml`, `compose.dev.yaml`: Compose 設定

## 作業時の注意
- 起動/停止はスクリプトを優先して使用する。
- Compose設定やスクリプトを変更した場合は、必要に応じて使い方ドキュメントも更新する。
- テストが必要な場合は、対象ディレクトリで実行する(この雛形には標準のテスト手順がない場合がある)。
