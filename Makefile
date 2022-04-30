.PHONY: migrate migrate_down migrate_up migrate_version docker prod docker_delve local swaggo test

# Version - this is optionally used on goto command
V?=

# Number of migrations - this is optionally used on up and down commands
N?=

# MySQL Credentials
MYSQL_USER ?= root
MYSQL_PASSWORD ?= root//14045
MYSQL_HOST ?= 127.0.0.1
MYSQL_DATABASE ?= samp_evxrp_dev
MYSQL_PORT ?= 3306

MYSQL_DSN ?= $(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

# go migrate
migrate_up:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path db/migrations up $(N)

migrate_down:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path db/migrations down $(N)

migrate_to_version:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path db/migrations goto $(V)

migrate_drop_db:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path db/migrations drop

migrate_force:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path db/migrations force $(V)

migration_version:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path db/migrations version


# docker-compose commands
develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.dev.yml up --build

prod:
	echo "Starting docker prod environment"
	docker-compose -f docker-compose.prod.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build


# tools commands
run-linter:
	echo "Starting linters"
	golangci-lint run ./...

# todo: go swagger
# swaggo:
# 	echo "Starting swagger generating"
# 	swag init -g **/**/*.go

# main
run:
	go run ./cmd/app/main.go

build:
	go build ./cmd/app/main.go

test:
	go test -cover ./...

# go modules
deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# docker supports
FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)
