.PHONY: proto postgres migrate_up migrate_down

proto:
	protoc -I ./proto --go_out=. --go-grpc_out=. proto/chat.proto

postgres:
	docker start postgresDB

migrate_up:
	migrate -source file://migrations \
			-database postgres://postgres:kjunn2000@localhost/straper_db?sslmode=disable up 

migrate_down:
	migrate -source file://migrations \
			-database postgres://postgres:kjunn2000@localhost/straper_db?sslmode=disable down
