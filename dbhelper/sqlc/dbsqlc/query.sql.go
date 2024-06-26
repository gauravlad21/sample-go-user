// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package dbsqlc

import (
	"context"
)

const deleteEmployee = `-- name: DeleteEmployee :exec
Delete from employee
WHERE id=$1
`

func (q *Queries) DeleteEmployee(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

const getEmployeeById = `-- name: GetEmployeeById :many
SELECT id, employee_name, position, salary, version, created, updated FROM employee
WHERE id=$1
`

func (q *Queries) GetEmployeeById(ctx context.Context, id int32) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, getEmployeeById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.EmployeeName,
			&i.Position,
			&i.Salary,
			&i.Version,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmployeeByPagination = `-- name: GetEmployeeByPagination :many
SELECT id, employee_name, position, salary, version, created, updated FROM employee
limit $1 offset $2
`

type GetEmployeeByPaginationParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetEmployeeByPagination(ctx context.Context, arg GetEmployeeByPaginationParams) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, getEmployeeByPagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.EmployeeName,
			&i.Position,
			&i.Salary,
			&i.Version,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertEmployee = `-- name: InsertEmployee :one
INSERT INTO employee(
    employee_name, position, salary
)VALUES(
    $1, $2, $3
)
RETURNING id
`

type InsertEmployeeParams struct {
	EmployeeName string
	Position     string
	Salary       float64
}

func (q *Queries) InsertEmployee(ctx context.Context, arg InsertEmployeeParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertEmployee, arg.EmployeeName, arg.Position, arg.Salary)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateEmployee = `-- name: UpdateEmployee :exec
UPDATE employee
SET employee_name=$1, position=$2, salary=$3, version=version+1
WHERE id=$4 and version=version
`

type UpdateEmployeeParams struct {
	EmployeeName string
	Position     string
	Salary       float64
	ID           int32
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee,
		arg.EmployeeName,
		arg.Position,
		arg.Salary,
		arg.ID,
	)
	return err
}
