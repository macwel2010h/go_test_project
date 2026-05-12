app-up : MySQL	
	sleep 7
	go mod tidy
	go run .

MySQL : Dockerfile docker-compose.yml database_init.sql
	open -a Docker
	sleep 2
	docker compose build mysql
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
