download migrate script
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

download sqlc binary
curl https://downloads.sqlc.dev/sqlc_1.27.0_linux_amd64.tar.gz | tar xvz

commands

init migrations structure
./migrate create -ext=sql -dir=sql/migrations -seq init

up migrations
./migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose up

down migrations
./migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose down
