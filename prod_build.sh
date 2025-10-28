#プロダクト環境初回ビルド時
docker compose -f docker-compose.yml -f docker-compose-prod.yml up -d --build