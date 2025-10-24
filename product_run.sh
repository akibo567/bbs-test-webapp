# 初回
#docker compose build
#2回目以降
#docker compose up -d
# ログ確認
#docker compose logs -f

#docker compose -f docker-compose.yml -f docker-compose-prod.yml up --watch


docker compose -f docker-compose.yml -f docker-compose-prod.yml up --build