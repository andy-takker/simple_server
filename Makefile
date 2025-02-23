local:
	docker compose -f docker-compose.dev.yaml up --build --force-recreate --remove-orphans

local-down:
	docker compose -f docker-compose.dev.yaml down -v
