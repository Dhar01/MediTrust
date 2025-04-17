// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package users

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	googleuuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Address defines model for Address.
type Address struct {
	City          *string `json:"city,omitempty"`
	Country       *string `json:"country,omitempty"`
	PostalCode    *string `json:"postal_code,omitempty"`
	StreetAddress *string `json:"street_address,omitempty"`
}

// Error500Problem defines model for Error500Problem.
type Error500Problem struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// FetchUserInfoResponse defines model for FetchUserInfoResponse.
type FetchUserInfoResponse struct {
	Address  *Address             `json:"address,omitempty"`
	Age      *int32               `json:"age,omitempty" validate:"gte=18"`
	Email    *openapi_types.Email `json:"email,omitempty"`
	IsActive *bool                `json:"is_active,omitempty"`
	Name     *FullName            `json:"name,omitempty"`
	Phone    *string              `json:"phone,omitempty"`
	Role     *string              `json:"role,omitempty"`
}

// FullName defines model for FullName.
type FullName struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

// UpdateUserRequest defines model for UpdateUserRequest.
type UpdateUserRequest struct {
	Address *Address  `json:"address,omitempty"`
	Age     *int32    `json:"age,omitempty" validate:"gte=18"`
	Name    *FullName `json:"name,omitempty"`
	Phone   *string   `json:"phone,omitempty"`
}

// UpdateUserResponse defines model for UpdateUserResponse.
type UpdateUserResponse struct {
	Address  Address             `json:"address"`
	Age      int32               `json:"age" validate:"gte=18"`
	Email    openapi_types.Email `json:"email"`
	IsActive bool                `json:"is_active"`
	Name     FullName            `json:"name"`
	Phone    string              `json:"phone"`
	Role     string              `json:"role"`
}

// User defines model for User.
type User struct {
	Address      *Address             `json:"address,omitempty"`
	Age          *int32               `json:"age,omitempty" validate:"gte=18"`
	Email        *openapi_types.Email `json:"email,omitempty"`
	HashPassword *string              `json:"hashPassword,omitempty"`
	Id           *googleuuid.UUID     `json:"id,omitempty"`
	IsActive     *bool                `json:"is_active,omitempty"`
	Name         *FullName            `json:"name,omitempty"`
	Phone        *string              `json:"phone,omitempty"`
	Role         *string              `json:"role,omitempty"`
}

// UserID defines model for UserID.
type UserID = googleuuid.UUID

// UpdateUserInfoByIDJSONRequestBody defines body for UpdateUserInfoByID for application/json ContentType.
type UpdateUserInfoByIDJSONRequestBody = UpdateUserRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete user data
	// (DELETE /users/{userID})
	DeleteUserByID(ctx echo.Context, userID UserID) error
	// Get user using userID
	// (GET /users/{userID})
	FetchUserInfoByID(ctx echo.Context, userID UserID) error
	// Update a user information using userID
	// (PUT /users/{userID})
	UpdateUserInfoByID(ctx echo.Context, userID UserID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", ctx.Param("userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUserByID(ctx, userID)
	return err
}

// FetchUserInfoByID converts echo context to params.
func (w *ServerInterfaceWrapper) FetchUserInfoByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", ctx.Param("userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FetchUserInfoByID(ctx, userID)
	return err
}

// UpdateUserInfoByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUserInfoByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userID" -------------
	var userID UserID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", ctx.Param("userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUserInfoByID(ctx, userID)
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

	router.DELETE(baseURL+"/users/:userID", wrapper.DeleteUserByID)
	router.GET(baseURL+"/users/:userID", wrapper.FetchUserInfoByID)
	router.PUT(baseURL+"/users/:userID", wrapper.UpdateUserInfoByID)

}

type InternalServerErrorResponse struct {
}

type InvalidInputResponse struct {
}

type NotFoundResponse struct {
}

type UnauthorizedAccessResponse struct {
}

type DeleteUserByIDRequestObject struct {
	UserID UserID `json:"userID"`
}

type DeleteUserByIDResponseObject interface {
	VisitDeleteUserByIDResponse(w http.ResponseWriter) error
}

type DeleteUserByID204Response struct {
}

func (response DeleteUserByID204Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteUserByID401Response = UnauthorizedAccessResponse

func (response DeleteUserByID401Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type FetchUserInfoByIDRequestObject struct {
	UserID UserID `json:"userID"`
}

type FetchUserInfoByIDResponseObject interface {
	VisitFetchUserInfoByIDResponse(w http.ResponseWriter) error
}

type FetchUserInfoByID200JSONResponse FetchUserInfoResponse

func (response FetchUserInfoByID200JSONResponse) VisitFetchUserInfoByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type FetchUserInfoByID400Response = InvalidInputResponse

func (response FetchUserInfoByID400Response) VisitFetchUserInfoByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type FetchUserInfoByID404Response = NotFoundResponse

func (response FetchUserInfoByID404Response) VisitFetchUserInfoByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type FetchUserInfoByID500Response = InternalServerErrorResponse

func (response FetchUserInfoByID500Response) VisitFetchUserInfoByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type UpdateUserInfoByIDRequestObject struct {
	UserID UserID `json:"userID"`
	Body   *UpdateUserInfoByIDJSONRequestBody
}

type UpdateUserInfoByIDResponseObject interface {
	VisitUpdateUserInfoByIDResponse(w http.ResponseWriter) error
}

type UpdateUserInfoByID202JSONResponse UpdateUserResponse

func (response UpdateUserInfoByID202JSONResponse) VisitUpdateUserInfoByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUserInfoByID400Response = InvalidInputResponse

func (response UpdateUserInfoByID400Response) VisitUpdateUserInfoByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UpdateUserInfoByID401Response = UnauthorizedAccessResponse

func (response UpdateUserInfoByID401Response) VisitUpdateUserInfoByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type UpdateUserInfoByID500Response = InternalServerErrorResponse

func (response UpdateUserInfoByID500Response) VisitUpdateUserInfoByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Delete user data
	// (DELETE /users/{userID})
	DeleteUserByID(ctx context.Context, request DeleteUserByIDRequestObject) (DeleteUserByIDResponseObject, error)
	// Get user using userID
	// (GET /users/{userID})
	FetchUserInfoByID(ctx context.Context, request FetchUserInfoByIDRequestObject) (FetchUserInfoByIDResponseObject, error)
	// Update a user information using userID
	// (PUT /users/{userID})
	UpdateUserInfoByID(ctx context.Context, request UpdateUserInfoByIDRequestObject) (UpdateUserInfoByIDResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteUserByID operation middleware
func (sh *strictHandler) DeleteUserByID(ctx echo.Context, userID UserID) error {
	var request DeleteUserByIDRequestObject

	request.UserID = userID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUserByID(ctx.Request().Context(), request.(DeleteUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUserByIDResponseObject); ok {
		return validResponse.VisitDeleteUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// FetchUserInfoByID operation middleware
func (sh *strictHandler) FetchUserInfoByID(ctx echo.Context, userID UserID) error {
	var request FetchUserInfoByIDRequestObject

	request.UserID = userID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.FetchUserInfoByID(ctx.Request().Context(), request.(FetchUserInfoByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "FetchUserInfoByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(FetchUserInfoByIDResponseObject); ok {
		return validResponse.VisitFetchUserInfoByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateUserInfoByID operation middleware
func (sh *strictHandler) UpdateUserInfoByID(ctx echo.Context, userID UserID) error {
	var request UpdateUserInfoByIDRequestObject

	request.UserID = userID

	var body UpdateUserInfoByIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateUserInfoByID(ctx.Request().Context(), request.(UpdateUserInfoByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateUserInfoByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateUserInfoByIDResponseObject); ok {
		return validResponse.VisitUpdateUserInfoByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY0W7buBL9FYI3j7IlO3aaGrjATZDrrtumCJoG+xAEwUQaS2wpUiWpNG6gf19wJNty",
	"pDTdbbfY7W5eYokzwzNzztAc3/NY54VWqJzls3tegIEcHRp6urBoFif+k1B8xgtwGQ+4ghz5jJf1YsAN",
	"fiyFwYTPnCkx4DbOMAfvtdQmB8dnPNU6lViWIuEBd6vC+1tnhEp5wO8GqR40L7eGw4sLir5ZHYi80Mb5",
	"uA2CnagEbsZT4bLyZhjrPKyXQ1qvqqrySG2hlUVKbqEcGgXyHM0tmv8bo41/naCNjSic0D7ltRGzZMWQ",
	"zKqAL9QtSJEsVFG6PjdaZYKWq4C/0W6uS5V0Ta0DV1qmtGNLsqgCfqGgdJk24jMmR3GM1nb92jYMaiOf",
	"Y119cjhKEtP4FkYXaJyoU4+FW/n/eAd5IX0pTzL4AB1uqoDHulTOPLA+BpVKSNBmfS6Ftg7kdawT3HUb",
	"jaNnfQ7WGUR3DVu4bZ99dgpCsXMy6rpXmzf65j3GVG0icxpFZ0bfSMy7BcjRWkgJXg8aT8gOimkUBVsx",
	"C+X2x1sgQjlM0fQjmaOLM2ojtdRvG/l18bRS3zO45DP+n3DbmGHDabgmtAp4g34DcT/qANrtrDXqu4GG",
	"Qgw8OymqAd45AwMHKW1OqgVHzeXwv6NDygpzEHKXFt/8/2sefbPxVn1q8x6ehb2G2InbXeT1odEY32gt",
	"EZS3rpv8ywWZl1K+8XZedplWDwQXjQ6i5o8HPBfqNarUHxOjUQ88o+UD/7i0TudUyq9Q3QZMh96lMNZd",
	"rxPaxn+pM7ULbL8Hl4Re5xONT/n2obwoPMFekm/xY4nW/d3U+IN18VQN/23pn6eltzeZyxprsMnckxNs",
	"SG0irxG263DVpxeL5p+qkAxsdgbWftIm2Q23N4a90XhvdfJs/tx+fj0/f3UXnb36JXmbvnhVDofDXr09",
	"iDGdRng4iaIBjp/fDCajZDKAZ6ODwWRycDCdTiaNTv5KF9Gf/HvQ36AwLo1wq/PtVfQYwaA5Kv2m9/yG",
	"nuZrUl7++o43F1fKnVa3e2XOFZzu70ItdfcmfHS2YEttmJcny0FBijkqx4RioJhWUihkRQYmh3jFAy5F",
	"jM2R3fB3unjnE3fCUbq+XdnpNtDR2YIH/BaNrTeMhtFw1O2qXCSJxE9gfHtf8tPW41XHmISUQ1EIVXch",
	"qeMx2dQa9CXQBSooBJ/x/WE0jBq9UYTQF8CG9/VgVtWFkuiwW7ITes9chkzqNEU/qVD9fNP5Mwq8oR/9",
	"mhC+JMcr6oPtcrKJ1FpuT5CX/ardmoTNhFldPRjOxtGkZ1QqacxZllKuWI0rIdR9J8UkGj3WNZudwp5B",
	"qy1gyqAt3csrj9SWeQ5+JlrXkZSXgKMJik7VSzosLb+qAp5iz4RIgwGD2tXr2veC0J4FoVK2Ga53yVi2",
	"x4lePuY9Ft+LkohmR60cKsoIikKKmPYO31uf1n1r+P/iedU7FlGP75bpvGxY8YRGTxO6M5aT0+Rpp81w",
	"XgV8+nW7dH86+F26eYGuZv4B213x9P68UN8/n1IPHYrdFg+Ygw9oWQHGCZCspGAdqZWbO+6jWrvoM/kG",
	"sdE4cqyT1XfTWXfYqXbvef5Lt+oIffynAHhc5V0ayS1h7RPvG3rgD56EP6obvk7OPe1R7+J3rcVWGtnc",
	"F2ZhKHUMMtPWzQ6jwyiEQoS3I15dVb8FAAD//+cLHiT0FAAA",
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
