Step to run this project :
- Run docker with this command "docker-compose -f docker-compose.yaml up" to run NSQ
- Create DB with name "db_source" & "db_destination"
- Run this command "make db-init-source && make db-init-destination" to migrate the table
- Run this command "go run seed/main.go" to seed the table
- Run this command "go run cmd/product-project-http/*.go" to run the http service
- Run this command "go run cmd/product-project-mq/*.go" to run the mq service

-- example request to update all product --
curl --location --request POST 'http://localhost:8000/update-all'