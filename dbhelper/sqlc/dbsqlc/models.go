// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package dbsqlc

import (
	"database/sql"
)

type Employee struct {
	ID           int32
	EmployeeName string
	Position     string
	Salary       float64
	Created      sql.NullTime
	Updated      sql.NullTime
}