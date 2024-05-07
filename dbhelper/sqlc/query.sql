-- name: InsertEmployee :one
INSERT INTO employee(
    employee_name, position, salary
)VALUES(
    $1, $2, $3
)
RETURNING id;

-- name: DeleteEmployee :exec
Delete from employee
WHERE id=$1;

-- name: UpdateEmployee :exec
UPDATE employee
SET employee_name=$1, position=$2, salary=$3, version=version+1
WHERE id=$4 and version=version;

-- name: GetEmployeeById :many
SELECT * FROM employee
WHERE id=$1;

-- name: GetEmployeeByPagination :many
SELECT * FROM employee
limit $1 offset $2;