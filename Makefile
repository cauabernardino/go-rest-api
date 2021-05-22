SHELL=/bin/bash

createmigration:
	migrate create -ext sql -dir db/migrations -seq create_products_table

postgres:
	source .env.dev && \
	docker run --name go_api_db -p 5432:5432 -e POSTGRES_USER=$${POSTGRES_USER} -e POSTGRES_PASSWORD=$${POSTGRES_PASSWORD} \
	-e POSTGRES_DB=$${POSTGRES_DB} -d postgres:13.2-alpine

migrateup:
	source .env.dev && \
	sleep 3 && \
	migrate -path db/migrations -database "$${POSTGRES_URL}" -verbose up

migratedown:
	source .env.dev && \
	migrate -path db/migrations -database "$${POSTGRES_URL}" -verbose down
	
testdbup: postgres migrateup

testdbdown:
	docker rm go_api_db -f

test: testdbup
	go test ./... -v

.PHONY: createmigration postgres migrateup migratedown testdbup testdbdown test