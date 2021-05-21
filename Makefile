createmigration:
	migrate create -ext sql -dir db/migrations -seq create_products_table



.PHONY: createmigration