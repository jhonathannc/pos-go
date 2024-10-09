createmigration:
	./migrate create -ext=sql -dir=sql/migrations -seq init

migrateup:
	./migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose up

migratedown:
	./migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose down

# ignorar arquivo migrate, ignorado aqui pra usar o arquivo bin localmente
# .PHONY: migrate
