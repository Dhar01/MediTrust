// Package product provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package product

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	googleuuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// MedicineRequest defines model for MedicineRequest.
type MedicineRequest = medicineRequest

// MedicineResponse defines model for MedicineResponse.
type MedicineResponse = medicineResponse

// MedicineID defines model for MedicineID.
type MedicineID = googleuuid.UUID

// CreateMedicineJSONRequestBody defines body for CreateMedicine for application/json ContentType.
type CreateMedicineJSONRequestBody = MedicineRequest

// UpdateMedicineInfoByIDJSONRequestBody defines body for UpdateMedicineInfoByID for application/json ContentType.
type UpdateMedicineInfoByIDJSONRequestBody = MedicineRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all medicines
	// (GET /medicines)
	FetchMedicineList(ctx echo.Context) error
	// Create a new medicine (admin only)
	// (POST /medicines)
	CreateMedicine(ctx echo.Context) error
	// Delete a medicine by ID (admin only)
	// (DELETE /medicines/{medicineID})
	DeleteMedicineByID(ctx echo.Context, medicineID MedicineID) error
	// Get a medicine by ID
	// (GET /medicines/{medicineID})
	FetchMedicineByID(ctx echo.Context, medicineID MedicineID) error
	// Update a medicine by ID (admin only)
	// (PUT /medicines/{medicineID})
	UpdateMedicineInfoByID(ctx echo.Context, medicineID MedicineID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FetchMedicineList converts echo context to params.
func (w *ServerInterfaceWrapper) FetchMedicineList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FetchMedicineList(ctx)
	return err
}

// CreateMedicine converts echo context to params.
func (w *ServerInterfaceWrapper) CreateMedicine(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateMedicine(ctx)
	return err
}

// DeleteMedicineByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteMedicineByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "medicineID" -------------
	var medicineID MedicineID

	err = runtime.BindStyledParameterWithOptions("simple", "medicineID", ctx.Param("medicineID"), &medicineID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medicineID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteMedicineByID(ctx, medicineID)
	return err
}

// FetchMedicineByID converts echo context to params.
func (w *ServerInterfaceWrapper) FetchMedicineByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "medicineID" -------------
	var medicineID MedicineID

	err = runtime.BindStyledParameterWithOptions("simple", "medicineID", ctx.Param("medicineID"), &medicineID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medicineID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FetchMedicineByID(ctx, medicineID)
	return err
}

// UpdateMedicineInfoByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateMedicineInfoByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "medicineID" -------------
	var medicineID MedicineID

	err = runtime.BindStyledParameterWithOptions("simple", "medicineID", ctx.Param("medicineID"), &medicineID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medicineID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateMedicineInfoByID(ctx, medicineID)
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

	router.GET(baseURL+"/medicines", wrapper.FetchMedicineList)
	router.POST(baseURL+"/medicines", wrapper.CreateMedicine)
	router.DELETE(baseURL+"/medicines/:medicineID", wrapper.DeleteMedicineByID)
	router.GET(baseURL+"/medicines/:medicineID", wrapper.FetchMedicineByID)
	router.PUT(baseURL+"/medicines/:medicineID", wrapper.UpdateMedicineInfoByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RX32/bNhD+VwhuDxsgW7Rrd62APSRLO3hIu2JtMGBBHmjqLLOlSJU8pRUC/e8DSdny",
	"DyVNu/5YHyUd73j3fffd6YYKU1ZGg0ZHsxtacctLQLDh6RnkUkgNizP/JDXNaMVxTROqeQk0o2VvkFAL",
	"b2tpIacZ2hoS6sQaSu5ProwtOdKMFsYUCupa5jSh2FTeh0MrdUET+n5UmFH3sjccX1wE79uvI1lWxqL3",
	"291iz2u4YEYLiet6ORamTOPnNHxv27b1N3WV0Q5Ckqc8/wve1uDwibXG+lc5OGFlhdL4lE95Tmy0SIjU",
	"11zJnEhd1UjbhC40gtVcvQR7DfYWFxsj4oIVgWDWJvS5waem1vkt554bJOE79deOBd1Dprt5gM6aCizK",
	"mNWenxsK73lZKV+sCwf+9gQtcCxBIzErUnCPgiC1Eh76Q2zahObG8QL2XU1ZWQzZllzXKy6wtmD3Tzxx",
	"b3gDK/JizW3JBdQoBVeOnGM+HvIUAd71cG4cCJRqyLqyUuybz5Oee1Ljg2l/TGqEAixNaACUo3+5oXBS",
	"IPzKAusMr+RImBwK0CN4j5aPkBeOZrpWyoOCRrzZCzph94l6h+ebO64U+Ns5M8vXIPCgc8oDZtzZOaXJ",
	"QQ02zdmaWzZJZUfctLImrwWm8YS/RE/B2EzHHBSeYpCf4CFtpvMRm4/Y41eTacZYxtg/dKdkPvERyhIG",
	"efgNeS3zfbv5nMGjGWMjmD5ejmaTfDbiv0wejmazhw/n89mMMcZ2ExvUve+tXz6Z8m1C6yr/MCEefAwh",
	"7t0MHUe/UDf4ooCorcTmpRfpbrIAt2BPau/thi7D09NNUn/8/Yp2ku5jx699gmvEKjqWemWOJ8PJiwVZ",
	"GUtKrnkhdUE2qTrfBFwTo5XUQKrInIYmVEkBXZ92CT9bvPKwoMSAw6ajycmLhddFsC4GY2M2nhxrVinz",
	"XME7bn26l/TZzuPVkXEoecmrSuqocaEbbhvUcer79E0FmleSZvTBmI1ZB0/wkG5z9k8F4HGZngKKNeFE",
	"SRcEgV9zqfhSQV8vGmJY7k8s8s2ZTS3OZdDQvY1hyliQN6MRdAjKq0pJEVykr12Upn79kQhlOPijhRXN",
	"6A9pv3Sl3VBPj+S0pza3ljeRDPvZnXdZ9bm0CZ3H2w3F2maRDm0tgcV1WXLb0Iz+Dki4Unt1iuPpkvbv",
	"rryMGDdQ+d+C+BNONLzbOiGg0TZj8qdWjScpz0upieCaVGB9vxNcS0e48E7GR8hEn5tKdSsnODw1efNR",
	"iNwPiDhAu41xd7dtjwgx+QLhNzw4xn3bqN2EJa4WApxb1Uo1ngOz+3DgcPH9z9zpFJBml/vad3nVXu1S",
	"a5gaP0UyGK2anwe4Rq98jL7l05v+56ON9FOAcExEC6W59tG2kVbWeKIBcWgsHLHsLDja1Pi0Cb8fuz9G",
	"l8MV6k3SnR8nn/oBWWbHl9wiGrP43hCNJdut8bIhi7MPYRr14w7h9rPPz0tpdFDv3n/t/NAL0Nyh3l8A",
	"PPZVO/1lpEFkwOzDSO7/T37maXCA7u0DoR4A9CLsfncgumyIRDcEaTy6hUWvzOfB9X8yO6ZflVF9B8Vl",
	"/LMpzbfg5331qSPfp+hTjOPjRpLVVnXreZamygiu1sZh9og9YimvZHo9oe1V+28AAAD//1fABvRYEwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
