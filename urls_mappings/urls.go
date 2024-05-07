package urlsmappings

import (
	"github.com/gauravlad21/sample-go-employee/controller"

	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type urlMap struct {
	Url     string
	Method  string
	Handler []gin.HandlerFunc
}

var urlsMappings []*urlMap

func GetUrlMaps() []*urlMap {
	return urlsMappings
}

func init() {
	urlsMappings = []*urlMap{
		// healthcheck endpoints
		{Url: "/hello", Method: GET, Handler: []gin.HandlerFunc{controller.Hello}},

		{Url: "/employee", Method: POST, Handler: []gin.HandlerFunc{controller.AddEmployee}},
		{Url: "/employee", Method: PUT, Handler: []gin.HandlerFunc{controller.UpdateEmployee}},
		{Url: "/employee", Method: DELETE, Handler: []gin.HandlerFunc{controller.DeleteEmployee}},
		{Url: "/employee", Method: GET, Handler: []gin.HandlerFunc{controller.GetEmployeeById}},
		{Url: "/employees", Method: GET, Handler: []gin.HandlerFunc{controller.GetEmployeeByPagiantion}},
	}
}
