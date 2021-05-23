SHELL=/bin/bash

createmigration:
	migrate create -ext sql -dir db/migrations -seq create_products_table

postgres:
	source .env.test && \
	docker run --name go_api_db -p 5432:5432 -e POSTGRES_USER=$${POSTGRES_USER} -e POSTGRES_PASSWORD=$${POSTGRES_PASSWORD} \
	-d postgres:13.2-alpine

postgresdown:
	docker rm go_api_db -f

# TEST ENVIRONMENT
createtestdb:
	source .env.test && \
	docker exec -it go_api_db createdb --username=$${POSTGRES_USER} --owner=$${POSTGRES_USER} $${POSTGRES_DB}

migrateup_test:
	source .env.test && \
	migrate -path db/migrations -database "$${POSTGRES_URL}" -verbose up

test_command:
	go test ./... -v

test: createtestdb migrateup_test test_command
	source .env.test && \
	docker exec -it go_api_db dropdb --username=$${POSTGRES_USER} $${POSTGRES_DB}


# DEV ENVIRONMENT
createdevdb:
	source .env.dev && \
	docker exec -it go_api_db createdb --username=$${POSTGRES_USER} --owner=$${POSTGRES_USER} $${POSTGRES_DB}

migrateup_dev:
	source .env.dev && \
	migrate -path db/migrations -database "$${POSTGRES_URL}" -verbose up

migratedown_dev:
	source .env.dev && \
	migrate -path db/migrations -database "$${POSTGRES_URL}" -verbose down

devdbup: createdevdb migrateup_dev

devdbdown:
	source .env.dev && \
	docker exec -it go_api_db dropdb $${POSTGRES_DB}


.PHONY: createmigration postgres postgresdown test devdbup devdbdown
