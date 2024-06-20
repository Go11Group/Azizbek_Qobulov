DB = postgres://postgres:root@localhost:5432/master

run:
	go run postman/main.go
migrate_file:
	migrate create -ext sql -dir migrations -seq nil
migrate_up:
	migrate -path migrations -database ${DB} -verbose up
migrate_down:
	migrate -path migrations -database ${DB} -verbose down
migrate_force:
	migrate -path migrations -database ${DB} -verbose force 1