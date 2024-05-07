package controller

import (
	"net/http"
	"strconv"

	"github.com/gauravlad21/sample-go-employee/common"
	"github.com/gin-gonic/gin"
)

func Hello(oldctx *gin.Context) {
	ctx := common.GetContext(oldctx)
	msg := serviceRepo.Hello(ctx)
	oldctx.JSON(200, msg)
}

func AddEmployee(c *gin.Context) {
	body := &common.Employee{}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.AddEmployee(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func UpdateEmployee(c *gin.Context) {
	body := &common.Employee{}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.UpdateEmployee(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func DeleteEmployee(c *gin.Context) {
	idStr := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read id"))
		return
	}
	res := serviceRepo.DeleteEmployee(common.GetContext(c), int32(id))
	c.JSON(http.StatusOK, res)
}

func GetEmployeeById(c *gin.Context) {
	idStr := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read id"))
		return
	}
	res := serviceRepo.GetEmployeeById(common.GetContext(c), int32(id))
	c.JSON(http.StatusOK, res)
}

func GetEmployeeByPagiantion(c *gin.Context) {
	offsetStr := c.Request.URL.Query().Get("offset")
	offset, errOffset := strconv.Atoi(offsetStr)
	LimitStr := c.Request.URL.Query().Get("limit")
	limit, errLimit := strconv.Atoi(LimitStr)

	if errOffset != nil || errLimit != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read offset or limit"))
		return
	}
	res := serviceRepo.GetEmployeeByPagination(common.GetContext(c),
		&common.PaginationReq{Limit: int32(limit), Offset: int32(offset)})
	c.JSON(http.StatusOK, res)
}
