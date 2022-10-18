package util

import (	
	"strconv"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Message *string				`json:"message"`
	Data	*interface{}		`json:"data"`
	Code	int					`json:"code"`
	Links 	Links				`json:"links"`
}

type Links struct {
	Self	string	`json:"self"`
	Next	string	`json:"next"`
	Last	string	`json:"last"`
}

func NewAPIResponse(data interface{}, message string, code int) APIResponse {
	response := APIResponse{
		Data	: &data,
		Message	: &message,
		Code	: code,
	}

	return response
}

func NewAPIPaginationResponse(data interface{}, message string, code int, ctx *gin.Context, pagination Pagination) APIResponse {
	
	links := NewLinks(ctx, pagination)

	response := APIResponse{
		Data	: &data,
		Message	: &message,
		Code	: code,
		Links	: links,
	}

	return response
}

func NewLinks(ctx *gin.Context, pagination Pagination) Links {

	selfUrl := ctx.Request.Host + ctx.Request.URL.Path + ctx.Request.URL.RawQuery

	var nextUrl string
	if ctx.Request.URL.Query().Has("page"){		
		q := ctx.Request.URL.Query()
		page, _ := strconv.Atoi(ctx.Request.URL.Query().Get("page"))
		if (page + 1) <= pagination.TotalPages  {
			page = page + 1
		} else {
			page = pagination.TotalPages
		}		
		q.Set("page", strconv.Itoa(page))	
		nextUrl = ctx.Request.Host + ctx.Request.URL.Path + q.Encode()		
	}

	var lastUrl string
	if ctx.Request.URL.Query().Has("page"){						
		q := ctx.Request.URL.Query()
		q.Set("page", strconv.Itoa(pagination.TotalPages))		
		lastUrl = ctx.Request.Host + ctx.Request.URL.Path + q.Encode()
	}
	
	return Links {
		Self: selfUrl,
		Next: nextUrl,
		Last: lastUrl,
	}
}