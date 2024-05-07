package common

type StatusCode int32

const (
	StatusCode_NO_STATUS_CODE StatusCode = 0
	StatusCode_OK             StatusCode = 200
	StatusCode_INTERNAL_ERROR StatusCode = 500
	StatusCode_BAD_REQUEST    StatusCode = 400
	StatusCode_CREATED        StatusCode = 201
)

type Response struct {
	StatusCode StatusCode `json:"status_code,omitempty"`
	Msg        string     `json:"msg,omitempty"`
	ErrorMsg   string     `json:"error,omitempty"`
}

func GetResponse(statusCode StatusCode, msgs string, errMsgs string) *Response {
	return &Response{
		StatusCode: statusCode,
		ErrorMsg:   errMsgs,
		Msg:        msgs,
	}
}

func GetDefaultResponse() *Response {
	return GetResponse(StatusCode_OK, "", "")
}

func GetErrMsgsResponse(statusCode StatusCode, errMsgs string) *Response {
	return GetResponse(statusCode, "", errMsgs)
}

func GetSuccessResponse(statusCode StatusCode, msgs string) *Response {
	return GetResponse(statusCode, msgs, "")
}

func GetErrResponse(statusCode StatusCode, err error) *Response {
	return GetResponse(statusCode, "", err.Error())
}
