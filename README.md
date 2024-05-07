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


generate mock for unit test:
```
mockgen -destination=unit_test/mocks/mock_dbops.go -package=mocks -source=dbhelper/dpOperations.go DbOperationsIF
```


usage of pessimistic locking [db trasaction] from service.go file
```
		tx, err := dbhelper.StartTransaction(ctx, dbhelper.GetDb())
		if err != nil {
			// error creating transaction
		}

		dbhelper.EndTransaction(ctx, tx, func(tx *sqlx.Tx) (txErr error) {
			err := s.DbOps.UpdateEmployee(ctx, req)
			if err != nil {
				return fmt.Errorf("UpdateEmployee::failed with error", err)
			}
			return nil
		}(tx))
```

usage optimistic locking  from query.go file
```
UPDATE employee
SET employee_name=$1, position=$2, salary=$3, version=version+1
WHERE id=$4 and version=version;
```