// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://t.me/Dhar01",
            "email": "loknathdhar66@gmail.com"
        },
        "license": {
            "name": "GPL v3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Authenticate a user with email and password to obtain an access token in the response body, while the refresh token is set as a secure HTTP-only cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "User log in request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LogIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Set-Cookie: refresh_token=\u003ctoken\u003e; HttpOnly; Secure; Path=/; Domain=\u003cyour-domain.com\u003e",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request received",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "logs out the authenticated user by invalidating the refresh token. The refresh token is cleared by setting an expired cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User Logout",
                "responses": {
                    "200": {
                        "description": "Set-Cookie: refresh_token=; HttpOnly; Secure; Path=/; Domain=\u003cyour-domain.com\u003e; Max-Age=0",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid or expired token",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medicines": {
            "get": {
                "description": "Fetch a list of available medicines",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "Get all medicines",
                "responses": {
                    "200": {
                        "description": "List of medicines",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Medicine"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new medicine on the store. Only an admin can create a medicine.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "Creates a medicine info - Admin only",
                "parameters": [
                    {
                        "description": "Create medicine details",
                        "name": "medicine",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateMedicineDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Medicine created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Medicine"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/medicines/{medID}": {
            "get": {
                "description": "Fetch information of a medicine by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "Get a medicine info by its ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Medicine ID",
                        "name": "medID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Medicine found",
                        "schema": {
                            "$ref": "#/definitions/models.Medicine"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates information of a medicine by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "Updates a medicine info by its ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Medicine ID",
                        "name": "medID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated medicine details",
                        "name": "medicine",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateMedicineDTO"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Medicine updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Medicine"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes information of a medicine by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicines"
                ],
                "summary": "Deletes a medicine info by its ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Medicine ID",
                        "name": "medID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Medicine deleted successfully",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves the refresh token from the cookie and generates a new access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Generate a new access token",
                "responses": {
                    "201": {
                        "description": "Access token generated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reset": {
            "post": {
                "description": "This endpoint resets the medicine, address, and user databases.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Reset all databases (development only)",
                "responses": {
                    "204": {
                        "description": "Database reset successfully"
                    },
                    "403": {
                        "description": "Forbidden – Not allowed outside development environment",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/revoke": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint revokes the refresh token, effectively logging them out.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Revoke refresh token",
                "responses": {
                    "204": {
                        "description": "Refresh token revoked successfully"
                    },
                    "401": {
                        "description": "Unauthorized – Invalid or missing refresh token",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Register a new user with email and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Sign up a user",
                "parameters": [
                    {
                        "description": "User signup request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ID: uuid",
                        "schema": {
                            "$ref": "#/definitions/models.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request received",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "to handler a user for admin, this handler will be used.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get user data by ID",
                "responses": {
                    "302": {
                        "description": "status found",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "bad request status",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "not found error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "updates user information for the logged in user, takes partial information update. request comes through isLoggedIn middleware.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User information update",
                "parameters": [
                    {
                        "description": "user update information request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "user update response data",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad request received",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid or expired token",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the logged in user, request comes through isLoggedIn middleware.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user data",
                "responses": {
                    "204": {
                        "description": "status no content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid or expired token",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "not found error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/reset": {
            "put": {
                "description": "user will submit update password and this handler will update the password on the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Updating user password",
                "responses": {
                    "202": {
                        "description": "status accepted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request received",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "if a user forget his/her password, they will request for password reset. A password reset link will be sent to the account email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Request for password reset",
                "responses": {
                    "202": {
                        "description": "status accepted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request sent",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/verify": {
            "get": {
                "description": "Upon signup, a autogenerated verify link will be sent to the user's email and this handler will verify that user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Verify a user on signup",
                "responses": {
                    "202": {
                        "description": "status accepted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request status",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Address": {
            "description": "Contains details of the user's location",
            "type": "object",
            "properties": {
                "city": {
                    "type": "string",
                    "example": "Dhaka"
                },
                "country": {
                    "type": "string",
                    "example": "Bangladesh"
                },
                "postal_code": {
                    "type": "string",
                    "example": "1207"
                },
                "street_address": {
                    "type": "string",
                    "example": "123 Main Street"
                }
            }
        },
        "models.CreateMedicineDTO": {
            "description": "DTO for creating a new medicine",
            "type": "object",
            "required": [
                "description",
                "dosage",
                "manufacturer",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Pain reliever"
                },
                "dosage": {
                    "type": "string",
                    "example": "500mg"
                },
                "manufacturer": {
                    "type": "string",
                    "example": "XZY Pharma"
                },
                "name": {
                    "type": "string",
                    "example": "Paracetamol"
                },
                "price": {
                    "type": "integer",
                    "format": "int32",
                    "minimum": 0,
                    "example": 50
                },
                "stock": {
                    "type": "integer",
                    "format": "int32",
                    "minimum": 0,
                    "example": 75
                }
            }
        },
        "models.ErrorResponse": {
            "description": "This struct represents the response structure for error handling.",
            "type": "object",
            "properties": {
                "code": {
                    "description": "HTTP status code",
                    "type": "integer",
                    "format": "int",
                    "example": 500
                },
                "message": {
                    "description": "Human-readable error message",
                    "type": "string",
                    "format": "string",
                    "example": "Internal server error"
                }
            }
        },
        "models.LogIn": {
            "description": "Contains credentials required for authentication",
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "SecurePass123"
                }
            }
        },
        "models.Medicine": {
            "description": "Medicine entity contains details about a medicine",
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Medicine creation time - auto-generated by the database",
                    "type": "string"
                },
                "description": {
                    "description": "Medicine Description",
                    "type": "string",
                    "example": "Pain reliever"
                },
                "dosage": {
                    "description": "Medicine Dosage",
                    "type": "string",
                    "example": "500mg"
                },
                "id": {
                    "description": "Unique ID of the user - auto generated by the database",
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "manufacturer": {
                    "description": "Medicine Manufacturer",
                    "type": "string",
                    "example": "XZY Pharma"
                },
                "name": {
                    "description": "Medicine Name",
                    "type": "string",
                    "example": "Paracetamol"
                },
                "price": {
                    "description": "Medicine price",
                    "type": "integer",
                    "format": "int32",
                    "example": 50
                },
                "stock": {
                    "description": "Medicine stock",
                    "type": "integer",
                    "format": "int32",
                    "example": 75
                },
                "updated_at": {
                    "description": "Medicine update time - auto-generated by the database",
                    "type": "string"
                }
            }
        },
        "models.Name": {
            "description": "Contains first and last name fields",
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string",
                    "minLength": 4,
                    "example": "John"
                },
                "lastname": {
                    "type": "string",
                    "minLength": 4,
                    "example": "Doe"
                }
            }
        },
        "models.ServerResponse": {
            "description": "Used for authentication responses",
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                }
            }
        },
        "models.SignUpResponse": {
            "description": "Contains the unique ID of the newly created user",
            "type": "object",
            "properties": {
                "id": {
                    "description": "unique ID of the user",
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "models.SignUpUser": {
            "description": "Contains required fields for creating a new user",
            "type": "object",
            "required": [
                "age",
                "email",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "minimum": 18,
                    "example": 25
                },
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "name": {
                    "$ref": "#/definitions/models.Name"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "SecurePass123"
                },
                "phone": {
                    "description": "for BD phone",
                    "type": "string",
                    "example": "01234567891"
                }
            }
        },
        "models.UpdateMedicineDTO": {
            "description": "DTO for updating a medicine information",
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Pain reliever"
                },
                "dosage": {
                    "type": "string",
                    "example": "500mg"
                },
                "manufacturer": {
                    "type": "string",
                    "example": "XZY Pharma"
                },
                "name": {
                    "type": "string",
                    "example": "Paracetamol"
                },
                "price": {
                    "type": "integer",
                    "format": "int32",
                    "example": 50
                },
                "stock": {
                    "type": "integer",
                    "format": "int32",
                    "example": 75
                }
            }
        },
        "models.UpdateUserDTO": {
            "description": "Contains optional fields for updating user information",
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/models.Address"
                },
                "age": {
                    "type": "integer",
                    "minimum": 18,
                    "example": 18
                },
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "name": {
                    "$ref": "#/definitions/models.Name"
                },
                "phone": {
                    "type": "string",
                    "example": "01234567891"
                }
            }
        },
        "models.UserResponseDTO": {
            "description": "Contains user profile information returned from API",
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/models.Address"
                },
                "age": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "$ref": "#/definitions/models.Name"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "refresh-token",
            "in": "cookie"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "MediTrust Backend API",
	Description:      "Documentation of api of online medicine pharmacy - MediTrust",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
