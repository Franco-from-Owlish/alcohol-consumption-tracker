.PHONY: start stop restart soft-restart
.PHONY: teardown nuke
.PHONY: build-server
.PHONY: make-migrations migrate

env/.env:
	cp env/sample.env env/.env

start: env/.env
	docker compose --env-file env/.env up -d --build

stop: env/.env
	docker compose --env-file env/.env down

restart: stop start

soft-restart: stop
	docker compose --env-file env/.env  up -d

teardown:
	docker compose down --rmi all

nuke: teardown
	rm -r docker/volumes || true
	rm env/.env || true

backend/bin/main:
	go build -o srb/bin/main backend/main.go

build-server: backend/bin/main

migrations:
	cd backend && atlas migrate diff --env gorm

migrate:
	cd backend && atlas migrate apply --env deploy