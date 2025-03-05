package dto

type ReqTokenDTO struct {
	RefreshToken string `json:"refresh_token"`
}

// ErrorResponse defines the structure of an error response
//
//	@description	This struct represents the response structure for error handling.
//	@example		{ "code": 500, "message": "Internal server error"}
type ErrorResponseDTO struct {
	Message string `json:"message" example:"Internal server error" format:"string"` // Human-readable error message
	Code    int    `json:"code" example:"500" format:"int"`                         // HTTP status code
}
