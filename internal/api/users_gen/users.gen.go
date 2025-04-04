// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Address defines model for Address.
type Address struct {
	City          *string `json:"city,omitempty"`
	Country       *string `json:"country,omitempty"`
	PostalCode    *string `json:"postal_code,omitempty"`
	StreetAddress *string `json:"street_address,omitempty"`
}

// FullName defines model for FullName.
type FullName struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

// SignUpRequest defines model for SignUpRequest.
type SignUpRequest struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// SignUpResponse defines model for SignUpResponse.
type SignUpResponse struct {
	UserId *openapi_types.UUID `json:"user_id,omitempty"`
}

// User defines model for User.
type User struct {
	Address      *Address             `json:"address,omitempty"`
	Age          *string              `json:"age,omitempty"`
	Email        *openapi_types.Email `json:"email,omitempty"`
	HashPassword *string              `json:"hashPassword,omitempty"`
	Id           *openapi_types.UUID  `json:"id,omitempty"`
	IsActive     *bool                `json:"is_active,omitempty"`
	Name         *FullName            `json:"name,omitempty"`
	Phone        *string              `json:"phone,omitempty"`
	Role         *string              `json:"role,omitempty"`
}

// UserSignUpHandlerJSONRequestBody defines body for UserSignUpHandler for application/json ContentType.
type UserSignUpHandlerJSONRequestBody = SignUpRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create/Sign Up a new user.
	// (POST /users)
	UserSignUpHandler(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// UserSignUpHandler operation middleware
func (siw *ServerInterfaceWrapper) UserSignUpHandler(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserSignUpHandler(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/users", wrapper.UserSignUpHandler)
}

type UserSignUpHandlerRequestObject struct {
	Body *UserSignUpHandlerJSONRequestBody
}

type UserSignUpHandlerResponseObject interface {
	VisitUserSignUpHandlerResponse(w http.ResponseWriter) error
}

type UserSignUpHandler201JSONResponse SignUpResponse

func (response UserSignUpHandler201JSONResponse) VisitUserSignUpHandlerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type UserSignUpHandler400Response struct {
}

func (response UserSignUpHandler400Response) VisitUserSignUpHandlerResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UserSignUpHandler500Response struct {
}

func (response UserSignUpHandler500Response) VisitUserSignUpHandlerResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Create/Sign Up a new user.
	// (POST /users)
	UserSignUpHandler(ctx context.Context, request UserSignUpHandlerRequestObject) (UserSignUpHandlerResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// UserSignUpHandler operation middleware
func (sh *strictHandler) UserSignUpHandler(ctx *gin.Context) {
	var request UserSignUpHandlerRequestObject

	var body UserSignUpHandlerJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UserSignUpHandler(ctx, request.(UserSignUpHandlerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UserSignUpHandler")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(UserSignUpHandlerResponseObject); ok {
		if err := validResponse.VisitUserSignUpHandlerResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
