package common

type Employee struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type PaginationReq struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type EmployeesResponse struct {
	CommonResponse *Response
	Employees      []*Employee
}
