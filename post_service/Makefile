CURRENT_DIR=$(shell pwd)

DB_URL="postgres://postgres:123@localhost:5432/ekzamen4db?sslmode=disable"

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-force:
	migrate -path migrations -database "$(DB_URL)" -verbose force 1

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_post_table




