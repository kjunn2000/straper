.PHONY : migrate_up migrate_down

migrate_up: 
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb up

migrate_action_up: 
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb up

migrate_down:
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb down

force_migrate_down:
		migrate -path db/migration \
			-database mysql://root:password@"(localhost:3306)"/straperdb force 0 

test:
		go test -v ./...

server:
		go run cmd/chat-server/main.go

mock:
		mockgen -package mock -destination pkg/storage/mock/store.go github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql Store 

build:
		docker build -t straper:latest .

docker-up:
		docker start chat-ws-mysql-1 straper-redis

docker-up-2:
		docker start chat-ws-mysql-1 straper-redis seaweedfs-master-1 seaweedfs-volume-1

docker-down:
		docker stop $(docker ps -q)