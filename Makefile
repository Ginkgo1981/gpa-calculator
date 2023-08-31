
.PHONY: dev
dev:
	go run app/gpa/gpa.go -f app/gpa/etc/gpa-api.dev.yaml

.PHONY: start
start:
	docker-compose --env-file deploy/.env up

.PHONY: create-db
create-db:
	docker exec -it ${CONTAINER_NAME} mysql -u"$(MYSQL_USER)" -p"$(MYSQL_PASSWORD)" -e "CREATE DATABASE IF NOT EXISTS gpa_dev"

.PHONY: migrate-up
migrate-up:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:23306)/gpa_dev" -path db/migrations -verbose up

.PHONY: import-seed-data
import-seed-data:
	docker exec -i ${CONTAINER_NAME} mysql -u"$(MYSQL_USER)" -p"$(MYSQL_PASSWORD)" gpa_dev < "$(SQL_FILE_PATH)"