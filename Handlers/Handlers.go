package handler

import (
	"context"

	"golang.org/x/tools/godoc/util"
)

type Login struct {
}

func (h *Login) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	} 
}

func chainMiddleware(route Route) []echo.MiddlewareFunc {
	// function for chain middleware implementation




	
}

func (h *Login) buildErrorResponse(ctx context.Context, response *pb.LoginResponse, c echo.Context, errorCode codes.Code, message string) (*util.Response, error) {

}
