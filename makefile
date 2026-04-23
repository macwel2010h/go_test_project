app-up : MySQL	
	sleep 10
	go mod tidy
	go run .

MySQL : Dockerfile docker-compose.yml database_init.sql
	docker compose up -d

reset-db:
	docker compose down
	rm -rf ./db
	@echo "Database wiped! Run 'make app-up' to rebuild and initialize."