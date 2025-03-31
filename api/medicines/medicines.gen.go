// Package medicines provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package medicines

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CreateMedicineDTO defines model for CreateMedicineDTO.
type CreateMedicineDTO struct {
	Description  string `json:"description"`
	Dosage       string `json:"dosage"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        string `json:"stock"`
}

// Medicine defines model for Medicine.
type Medicine struct {
	Description  *string             `json:"description,omitempty"`
	Dosage       *string             `json:"dosage,omitempty"`
	Id           *openapi_types.UUID `json:"id,omitempty"`
	Manufacturer *string             `json:"manufacturer,omitempty"`
	Name         *string             `json:"name,omitempty"`
	Price        *int                `json:"price,omitempty"`
	Stock        *int                `json:"stock,omitempty"`
}

// UpdateMedicineDTO defines model for UpdateMedicineDTO.
type UpdateMedicineDTO struct {
	Description  *string `json:"description,omitempty"`
	Dosage       *string `json:"dosage,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Name         *string `json:"name,omitempty"`
	Price        *int    `json:"price,omitempty"`
	Stock        *int    `json:"stock,omitempty"`
}

// MedicineCreateHandlerJSONRequestBody defines body for MedicineCreateHandler for application/json ContentType.
type MedicineCreateHandlerJSONRequestBody = CreateMedicineDTO

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new medicine (admin only)
	// (POST /medicines)
	MedicineCreateHandler(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// MedicineCreateHandler operation middleware
func (siw *ServerInterfaceWrapper) MedicineCreateHandler(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.MedicineCreateHandler(c)
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

	router.POST(options.BaseURL+"/medicines", wrapper.MedicineCreateHandler)
}

type MedicineCreateHandlerRequestObject struct {
	Body *MedicineCreateHandlerJSONRequestBody
}

type MedicineCreateHandlerResponseObject interface {
	VisitMedicineCreateHandlerResponse(w http.ResponseWriter) error
}

type MedicineCreateHandler201JSONResponse Medicine

func (response MedicineCreateHandler201JSONResponse) VisitMedicineCreateHandlerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type MedicineCreateHandler400Response struct {
}

func (response MedicineCreateHandler400Response) VisitMedicineCreateHandlerResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type MedicineCreateHandler500Response struct {
}

func (response MedicineCreateHandler500Response) VisitMedicineCreateHandlerResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Create a new medicine (admin only)
	// (POST /medicines)
	MedicineCreateHandler(ctx context.Context, request MedicineCreateHandlerRequestObject) (MedicineCreateHandlerResponseObject, error)
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

// MedicineCreateHandler operation middleware
func (sh *strictHandler) MedicineCreateHandler(ctx *gin.Context) {
	var request MedicineCreateHandlerRequestObject

	var body MedicineCreateHandlerJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.MedicineCreateHandler(ctx, request.(MedicineCreateHandlerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MedicineCreateHandler")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(MedicineCreateHandlerResponseObject); ok {
		if err := validResponse.VisitMedicineCreateHandlerResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
