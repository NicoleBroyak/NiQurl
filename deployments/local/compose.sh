docker compose --env-file ./../../config/niqurlconfigs/envfile.env down
docker compose --env-file ./../../config/niqurlconfigs/envfile.env build --no-cache
docker compose --env-file ./../../config/niqurlconfigs/envfile.env up -d