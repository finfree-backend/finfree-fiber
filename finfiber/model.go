package finfiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"time"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

func NewErrorResponse(err *fiber.Error) *ErrorResponse {
	return &ErrorResponse{
		Code:    err.Code,
		Message: err.Message,
		Time:    time.Now().Unix(),
	}
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Time    int64       `json:"time"`
	NextUrl string      `json:"next_url,omitempty"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Status: "OK",
		Time:   time.Now().Unix(),
		Data:   data,
	}
}

// data: response body
// URI: URI of request ( can be reached by ctx.Request().URI() )
// totalCount: total count of documents
func NewSuccessResponseWithNextUrl(data interface{}, URI *fasthttp.URI, totalCount int) *SuccessResponse {
	resp := NewSuccessResponse(data)

	page, err := URI.QueryArgs().GetUint(PAGE_QUERY_KEY)
	if err != nil {
		return resp
	}

	size, err := URI.QueryArgs().GetUint(SIZE_QUERY_KEY)
	if err != nil {
		return resp
	}

	if (page+1)*size < totalCount {
		URI.QueryArgs().Set(PAGE_QUERY_KEY, fmt.Sprintf("%d", page+1))
		resp.NextUrl = URI.String()
	}
	return resp
}
