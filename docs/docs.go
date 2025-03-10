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
            "email": "me@rizkirmdhn.cloud"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User Login Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_UserLoginResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "/auth/register": {
            "post": {
                "description": "Register a new user account with a unique email address.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User Registration Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_UserRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error or email already taken",
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
        "/options/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new option",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Options"
                ],
                "summary": "Create a new option",
                "parameters": [
                    {
                        "description": "Option data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateOption"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Option created",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_CreateOptionsResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/options/{id}/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete an option",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Options"
                ],
                "summary": "Delete an option",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Option ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Option deleted",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Option not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/options/{id}/update": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update an option",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Options"
                ],
                "summary": "Update an option",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Option ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Option data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateOption"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Option updated",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_CreateOptionsResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Option not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/polls": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all polls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Get all polls",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-array_models_PollsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "/polls/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create new poll",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Store new poll",
                "parameters": [
                    {
                        "description": "Poll Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreatePolling"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_CreatePollingResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error or title already taken",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "/polls/public": {
            "get": {
                "description": "Get public polls, options for voting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Get public polls, options for voting",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Poll Code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Polls found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_PublicPollsResponse"
                        }
                    },
                    "404": {
                        "description": "Poll not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/polls/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Show poll",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Show poll",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Poll ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_PollsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Poll not found",
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
        "/polls/{id}/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete poll",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Delete poll",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Poll ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithMessage"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Poll not found",
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
        "/polls/{id}/options": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all options of a poll",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Get all options of a poll",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Poll ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Options found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_CreateOptionsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Poll not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/polls/{id}/update": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update poll",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Polls"
                ],
                "summary": "Update poll",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Poll ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Poll Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdatePolling"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseWithData-models_UpdatePollingResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error or title already taken",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Poll not found",
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
        }
    },
    "definitions": {
        "models.CreateOptionsResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "votes_count": {
                    "type": "integer"
                }
            }
        },
        "models.CreatePollingResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PollsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.PublicPollsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CreateOptionsResponse"
                    }
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-array_models_PollsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PollsResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_CreateOptionsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CreateOptionsResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_CreatePollingResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CreatePollingResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_PollsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.PollsResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_PublicPollsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.PublicPollsResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_UpdatePollingResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.UpdatePollingResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_UserLoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.UserLoginResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithData-models_UserRegisterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.UserRegisterResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseWithMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Status": {
            "type": "string",
            "enum": [
                "active",
                "done"
            ],
            "x-enum-varnames": [
                "Active",
                "Done"
            ]
        },
        "models.UpdatePollingResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.UserLoginResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.UserRegisterResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "requests.CreateOption": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "poll_id": {
                    "type": "string"
                }
            }
        },
        "requests.CreatePolling": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string",
                    "format": "date-time",
                    "example": "2022-01-01 00:00"
                },
                "start_date": {
                    "type": "string",
                    "format": "date-time",
                    "example": "2022-01-01 00:00"
                },
                "status": {
                    "description": "testing:\n* active - Active, can be voted\n* done - Done, can't be voted",
                    "type": "string",
                    "enum": [
                        "active",
                        "done"
                    ]
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateOption": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "poll_id": {
                    "type": "string"
                }
            }
        },
        "requests.UpdatePolling": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "requests.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.UserRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "evote-be API",
	Description:      "This is a sample server evote-be server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
