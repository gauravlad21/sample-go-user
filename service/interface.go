package service

import (
	"context"

	"github.com/gauravlad21/sample-go-employee/common"
	"github.com/gauravlad21/sample-go-employee/dbhelper"
)

type ServiceIF interface {
	Hello(ctx context.Context) string
	AddEmployee(ctx context.Context, product *common.Employee) *common.Response
	UpdateEmployee(ctx context.Context, req *common.Employee) *common.Response
	DeleteEmployee(ctx context.Context, id int32) *common.Response
	GetEmployeeById(ctx context.Context, id int32) *common.EmployeesResponse
	GetEmployeeByPagination(ctx context.Context, req *common.PaginationReq) *common.EmployeesResponse
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
}

func New(dbOps dbhelper.DbOperationsIF) ServiceIF {
	return &ServiceStruct{
		DbOps: dbOps,
	}
}
