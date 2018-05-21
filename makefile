DB_NAME=mybitly

test:
	go test -v ./...

serve:
	go build
	go run main.go

clean_db:
	psql $(DB_NAME) < db/scripts/drop_tables.sql

bootstrap_db:
	psql < db/scripts/create_database.sql 
	psql $(DB_NAME) < db/scripts/create_short_link.sql 
	psql $(DB_NAME) < db/scripts/create_short_link_visit.sql 
