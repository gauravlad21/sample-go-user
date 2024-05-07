package dbhelper

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gauravlad21/sample-go-employee/common"
	"github.com/gauravlad21/sample-go-employee/dbhelper/sqlc/dbsqlc"
	_ "github.com/lib/pq"
)

type DbOperationsIF interface {
	Exec(ctx context.Context, query string) error
	CloseDb(ctx context.Context) error

	InsertEmployee(ctx context.Context, req *common.Employee, tx ...*sql.Tx) (id int32, err error)
	DeleteEmployee(ctx context.Context, id int32, tx ...*sql.Tx) (err error)
	UpdateEmployee(ctx context.Context, req *common.Employee, tx ...*sql.Tx) (err error)
	GetEmployeeById(ctx context.Context, id int32, tx ...*sql.Tx) (emp *common.Employee, err error)
	GetEmployeeByPagination(ctx context.Context, offset int32, limit int32, tx ...*sql.Tx) (emp []*common.Employee, err error)
}

type DbOps struct {
	DB     *sql.DB
	DbSqlc *dbsqlc.Queries
}

func New(db *sql.DB) DbOperationsIF {
	return &DbOps{DbSqlc: dbsqlc.New(db), DB: db}
}

func (dbOps *DbOps) Exec(ctx context.Context, query string) error {
	_, err := dbOps.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) CloseDb(ctx context.Context) error {
	return dbOps.DB.Close()
}

// start from below
func (dbOps *DbOps) InsertEmployee(ctx context.Context, req *common.Employee, tx ...*sql.Tx) (id int32, err error) {
	params := dbsqlc.InsertEmployeeParams{EmployeeName: req.Name, Position: req.Position, Salary: req.Salary}
	id, err = GetSqlcQuery(dbOps.DbSqlc, tx...).InsertEmployee(ctx, params)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dbOps *DbOps) DeleteEmployee(ctx context.Context, id int32, tx ...*sql.Tx) (err error) {
	err = GetSqlcQuery(dbOps.DbSqlc, tx...).DeleteEmployee(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) UpdateEmployee(ctx context.Context, req *common.Employee, tx ...*sql.Tx) (err error) {
	args := dbsqlc.UpdateEmployeeParams{ID: req.ID, EmployeeName: req.Name, Position: req.Position, Salary: req.Salary}
	err = GetSqlcQuery(dbOps.DbSqlc, tx...).UpdateEmployee(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) GetEmployeeById(ctx context.Context, id int32, tx ...*sql.Tx) (employee *common.Employee, err error) {
	employees, err := GetSqlcQuery(dbOps.DbSqlc, tx...).GetEmployeeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(employees) == 0 {
		return nil, fmt.Errorf("no employee found with id: %v", id)
	}
	emp := employees[0]
	employee = &common.Employee{
		ID:       emp.ID,
		Name:     emp.EmployeeName,
		Position: emp.Position,
		Salary:   emp.Salary,
	}
	return employee, nil
}

func (dbOps *DbOps) GetEmployeeByPagination(ctx context.Context, offset int32, limit int32, tx ...*sql.Tx) (allEmployees []*common.Employee, err error) {
	params := dbsqlc.GetEmployeeByPaginationParams{Limit: limit, Offset: offset}
	employees, err := GetSqlcQuery(dbOps.DbSqlc, tx...).GetEmployeeByPagination(ctx, params)
	if err != nil {
		return nil, err
	}
	for _, emp := range employees {
		allEmployees = append(allEmployees, &common.Employee{
			ID:       emp.ID,
			Name:     emp.EmployeeName,
			Position: emp.Position,
			Salary:   emp.Salary,
		})
	}
	return allEmployees, nil
}
