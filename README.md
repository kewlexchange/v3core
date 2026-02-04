# v3core
KEWL v3 Core


EXECUTE

go run .

TEST:

core/workers/exchange/dexv2/scanner

go test ./... -v    



# DOCKER BUID

docker build -t corev3 .
docker run -d --name corev3 -p 8080:8080 corev3

docker start corev3

# DOCKER CONTAINERS

docker ps

## TERMINAL

docker exec -it dd7b04567110 sh

## GO CONSOLE LOGS

docker logs -f dd7b04567110

## DOCKER IGNORE

echo ".env" >> .dockerignore



docker-compose down
docker-compose up -d --build