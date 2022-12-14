# run backend service

# source
db-init-source:
	go run migration/source/main.go init
	go run migration/source/main.go up

db-up-source:
	go run migration/source/main.go up

db-down-source:
	go run migration/source/main.go down

db-reset-source:
	go run migration/source/main.go reset

# destination
db-init-destination:
	go run migration/destination/main.go init
	go run migration/destination/main.go up

db-up-destination:
	go run migration/destination/main.go up

db-down-destination:
	go run migration/destination/main.go down

db-reset-destination:
	go run migration/destination/main.go reset
