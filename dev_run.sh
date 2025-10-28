#docker compose -f docker-compose.yml -f compose.dev.yaml up --watch
# または、ログを分けたいなら
#docker compose -f docker-compose.yml -f compose.dev.yaml watch


#docker compose -f docker-compose.yml -f docker-compose-dev.yml up --watch

docker compose -f docker-compose.yml -f docker-compose-dev.yml start
