mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest
	
migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/diary" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/diary" -verbose down
sqlc:
	sqlc generate

.PHONY: mysql migrateup migratedown sqlc