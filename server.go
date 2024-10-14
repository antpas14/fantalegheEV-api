// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Rank defines model for rank.
type Rank struct {
	EvPoints *float64 `json:"evPoints,omitempty"`
	Points   *int     `json:"points,omitempty"`
	Team     *string  `json:"team,omitempty"`
}

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse struct {
	Message *string `json:"message,omitempty"`
}

// ServerErrorResponse defines model for ServerErrorResponse.
type ServerErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// SuccessResponse defines model for SuccessResponse.
type SuccessResponse = []Rank

// GetCalculateParams defines parameters for GetCalculate.
type GetCalculateParams struct {
	LeagueName string `form:"leagueName" json:"leagueName"`
}

// CalculateJSONBody defines parameters for Calculate.
type CalculateJSONBody struct {
	Cookies    string `json:"cookies"`
	LeagueName string `json:"leagueName"`
}

// CalculateJSONRequestBody defines body for Calculate for application/json ContentType.
type CalculateJSONRequestBody CalculateJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get calculation
	// (GET /calculate)
	GetCalculate(ctx echo.Context, params GetCalculateParams) error
	// Calculate
	// (POST /calculate)
	Calculate(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCalculate converts echo context to params.
func (w *ServerInterfaceWrapper) GetCalculate(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCalculateParams
	// ------------- Required query parameter "leagueName" -------------

	err = runtime.BindQueryParameter("form", true, true, "leagueName", ctx.QueryParams(), &params.LeagueName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter leagueName: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCalculate(ctx, params)
	return err
}

// Calculate converts echo context to params.
func (w *ServerInterfaceWrapper) Calculate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Calculate(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/calculate", wrapper.GetCalculate)
	router.POST(baseURL+"/calculate", wrapper.Calculate)

}
