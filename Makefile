up:
	docker compose -f deployments/compose.yaml up -d --build
dev-up:
	docker compose -f deployments/compose.yaml up -d --build
down:
	docker compose -f deployments/compose.yaml down