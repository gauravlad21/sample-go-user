# sample-go-user
go-skeleton code for simple CRUD APIs.

run steps:
```
go mod tidy
go mod vendor
go build -o main .
./main  --config=./config.json
```
==========================================

insert employee
```
curl --location 'http://localhost:8002/employee?id=1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "emp 4",
    "position": "sde 2",
    "salary": 400.0
}'
```

update employee
```
curl --location --request PUT 'http://localhost:8002/employee' \
--header 'Content-Type: application/json' \
--data '{
    "id": 4, 
    "name": "emp 4",
    "position": "sde 2",
    "salary": 450.0
}'
```

delete employee
```
curl --location --request DELETE 'http://localhost:8002/employee?id=4'
```

get employee by id
```
curl --location 'http://localhost:8002/employee?id=1'
```

get employee with pagination
```
curl --location 'http://localhost:8002/employees?limit=2&offset=2'
```