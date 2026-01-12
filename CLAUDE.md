# AGENTS.md
このリポジトリで作業するAIエージェント向けの最小ガイドです。

## 概要
- Next.js(Frontend) + Go(Backend) + PostgreSQL + nginx を Docker Compose で動かしている
- 開発/確認環境はシェルスクリプト経由で起動する前提

## 主要ディレクトリ
- `frontend/`: Next.js フロントエンド
- `api/`: Go バックエンド
- `db/`: DB関連(初期化スクリプト等)
- `nginx/`: nginx 設定
- `docker-compose*.yml`: Compose 設定
## 作業時の注意
- Compose設定やスクリプトを変更した場合は、必要に応じて使い方ドキュメントも更新する。
