up:
	docker compose --env-file .env up --build -d
	scripts/grafana_setup.sh

down:
	docker compose --env-file .env down --rmi all --volumes --remove-orphans
