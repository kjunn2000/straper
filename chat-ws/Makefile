.PHONY : migrate_up migrate_down

migrate_up: 
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb up

migrate_actions_up:
		migrate -path db/migration \
			-database mysql://root:root@"(localhost:3306)"/straperdb up

migrate_down:
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb down

force_migrate_down:
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb force 0 

test:
		go test -v ./...