package service

import (
	"context"
	"fmt"

	"github.com/gauravlad21/sample-go-employee/common"
)

func (s *ServiceStruct) Hello(ctx context.Context) string {
	return "hello from service"
}

func (s *ServiceStruct) AddEmployee(ctx context.Context, req *common.Employee) *common.Response {

	// validation
	if req == nil || req.Name == "" || req.Position == "" || req.Salary == 0 {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "body is nil or input invalid")
	}

	id, err := s.DbOps.InsertEmployee(ctx, req)
	if err != nil {
		return &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: err.Error()}
	}
	return common.GetResponse(common.StatusCode_OK, fmt.Sprintf("employee created with ID: %v", id), "")
}

func (s *ServiceStruct) UpdateEmployee(ctx context.Context, req *common.Employee) *common.Response {
	// validation
	if req == nil || req.ID == 0 || req.Name == "" || req.Position == "" || req.Salary == 0 {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "input is nil")
	}

	err := s.DbOps.UpdateEmployee(ctx, req)
	if err != nil {
		return common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)
	}
	return common.GetDefaultResponse()
}

func (s *ServiceStruct) DeleteEmployee(ctx context.Context, id int32) *common.Response {
	err := s.DbOps.DeleteEmployee(ctx, id)
	if err != nil {
		return common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)
	}
	return common.GetDefaultResponse()
}

func (s *ServiceStruct) GetEmployeeById(ctx context.Context, id int32) *common.EmployeesResponse {
	employee, err := s.DbOps.GetEmployeeById(ctx, id)
	if err != nil || employee == nil {
		return &common.EmployeesResponse{Status: common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)}
	}
	return &common.EmployeesResponse{Status: common.GetDefaultResponse(), Employees: []*common.Employee{employee}}
}

func (s *ServiceStruct) GetEmployeeByPagination(ctx context.Context, req *common.PaginationReq) *common.EmployeesResponse {

	if req.Limit > common.MAX_ALLOW_LIMIT {
		return nil
	}

	employees, err := s.DbOps.GetEmployeeByPagination(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil
	}
	return &common.EmployeesResponse{Status: common.GetDefaultResponse(), Employees: employees}
}
