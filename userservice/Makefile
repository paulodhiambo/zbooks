SHELL := /bin/bash

include .env

run:
	go run cmd/userservice/main.go

migrate:
	migrate -path migrations -database "$(DB_CONNECTION_STRING)" up

make_migration:
	migrate create -dir migrations -ext sql -seq -digits 3 $(message)

proto-user:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user/*.proto

proto-role:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/role/*.proto

proto-permission:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/permission/*.proto

proto-gen: proto-user proto-role proto-permission

.PHONY: proto-gen run migrate make_migration