package servicetest

import (
	"context"
	"fmt"
	"testing"

	"github.com/gauravlad21/sample-go-employee/common"
	"github.com/gauravlad21/sample-go-employee/service"
	"github.com/gauravlad21/sample-go-employee/unit_test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/zeebo/assert"
)

func TestInsertEmployeeSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	var dbId int32 = 123
	m.EXPECT().InsertEmployee(gomock.Any(), gomock.Any()).Return(dbId, nil)

	svc := service.New(m)

	employee := &common.Employee{Name: "name", Position: "position", Salary: 10}

	res := svc.AddEmployee(context.Background(), employee)
	assert.Equal(t, res.StatusCode, common.StatusCode_OK)
	assert.Equal(t, res.Msg, fmt.Sprintf("employee created with ID: %v", dbId))
}

func TestInsertEmployeeFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	svc := service.New(m)

	employee := &common.Employee{Name: "name"}

	res := svc.AddEmployee(context.Background(), employee)
	assert.Equal(t, res.StatusCode, common.StatusCode_BAD_REQUEST)
	assert.Equal(t, res.ErrorMsg, "body is nil or input invalid")
}

func TestUpdateEmployeeSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	m.EXPECT().UpdateEmployee(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.New(m)

	employee := &common.Employee{ID: 123, Name: "name", Position: "position", Salary: 10}
	res := svc.UpdateEmployee(context.Background(), employee)
	assert.Equal(t, res.StatusCode, common.StatusCode_OK)
}

func TestUpdateEmployeeFail1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	svc := service.New(m)

	employee := &common.Employee{Name: "name", Position: "position", Salary: 10}
	res := svc.UpdateEmployee(context.Background(), employee)
	assert.Equal(t, res.StatusCode, common.StatusCode_BAD_REQUEST)
	assert.Equal(t, res.ErrorMsg, "input is nil")
}

func TestUpdateEmployeeFail2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	errMsg := "some error"
	m.EXPECT().UpdateEmployee(gomock.Any(), gomock.Any()).Return(fmt.Errorf(errMsg))

	svc := service.New(m)

	employee := &common.Employee{ID: 123, Name: "name", Position: "position", Salary: 10}
	res := svc.UpdateEmployee(context.Background(), employee)
	assert.Equal(t, res.StatusCode, common.StatusCode_INTERNAL_ERROR)
	assert.Equal(t, res.ErrorMsg, errMsg)
}

func TestDeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	m.EXPECT().DeleteEmployee(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.New(m)

	var id int32 = 10
	res := svc.DeleteEmployee(context.Background(), id)
	assert.Equal(t, res.StatusCode, common.StatusCode_OK)
}

func TestDeleteEmployeeFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	errMsg := "id not found"
	m.EXPECT().DeleteEmployee(gomock.Any(), gomock.Any()).Return(fmt.Errorf(errMsg))

	svc := service.New(m)

	var id int32 = 10
	res := svc.DeleteEmployee(context.Background(), id)
	assert.Equal(t, res.StatusCode, common.StatusCode_INTERNAL_ERROR)
	assert.Equal(t, res.ErrorMsg, errMsg)
}

func TestGetEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	emp := &common.Employee{Name: "name", Position: "position", Salary: 123.45, ID: 123}
	m.EXPECT().GetEmployeeById(gomock.Any(), gomock.Any()).Return(emp, nil)

	svc := service.New(m)

	var id int32 = 10
	res := svc.GetEmployeeById(context.Background(), id)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Status)
	assert.Equal(t, res.Status.StatusCode, common.StatusCode_OK)
	assert.NotNil(t, res.Employees)

	assert.Equal(t, len(res.Employees), 1)
	assert.Equal(t, res.Employees[0].Name, emp.Name)
	assert.Equal(t, res.Employees[0].Position, emp.Position)
	assert.Equal(t, res.Employees[0].Salary, emp.Salary)
}

func TestGetEmployeeFail1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	errMsg := "no emp found for id"
	m.EXPECT().GetEmployeeById(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf(errMsg))

	svc := service.New(m)

	var id int32 = 10
	res := svc.GetEmployeeById(context.Background(), id)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Status)
	assert.Equal(t, res.Status.StatusCode, common.StatusCode_INTERNAL_ERROR)
	assert.Equal(t, res.Status.ErrorMsg, errMsg)
	assert.Nil(t, res.Employees)
}

func TestGetEmployeesByPagination(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	emp := []*common.Employee{
		{Name: "name", Position: "position", Salary: 123.45, ID: 123},
		{Name: "name2", Position: "position2", Salary: 456.45, ID: 890},
	}
	m.EXPECT().GetEmployeeByPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(emp, nil)

	svc := service.New(m)

	req := &common.PaginationReq{Limit: 2, Offset: 0}
	res := svc.GetEmployeeByPagination(context.Background(), req)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Status)
	assert.Equal(t, res.Status.StatusCode, common.StatusCode_OK)
	assert.NotNil(t, res.Employees)

	assert.Equal(t, len(res.Employees), 2)
	assert.Equal(t, res.Employees[0].Name, emp[0].Name)
	assert.Equal(t, res.Employees[0].Position, emp[0].Position)
	assert.Equal(t, res.Employees[0].Salary, emp[0].Salary)

	assert.Equal(t, res.Employees[1].Name, emp[1].Name)
	assert.Equal(t, res.Employees[1].Position, emp[1].Position)
	assert.Equal(t, res.Employees[1].Salary, emp[1].Salary)
}

func TestGetEmployeesByPaginationFail1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockDbOperationsIF(ctrl)

	errMsg := "no emp found for id"
	m.EXPECT().GetEmployeeByPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf(errMsg))

	svc := service.New(m)

	req := &common.PaginationReq{Limit: 2, Offset: 0}
	res := svc.GetEmployeeByPagination(context.Background(), req)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Status)
	assert.Equal(t, res.Status.StatusCode, common.StatusCode_INTERNAL_ERROR)
	assert.Equal(t, res.Status.ErrorMsg, errMsg)
	assert.Nil(t, res.Employees)
}
