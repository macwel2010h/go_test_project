app-up : MySQL	
	sleep 10
	go mod tidy
	go run .

MySQL : Dockerfile docker-compose.yml database_init.sql
	docker compose up -d

reset-db:
	docker compose down
	rm -rf ./db
	docker rmi go_test_mysql_image:latest
	echo "Database wiped! Run 'make app-up' to rebuild and initialize."

reset-docker:
	docker stop go_test_mysql_container
	docker rm go_test_mysql_container
	docker rmi go_test_mysql_image:latest
