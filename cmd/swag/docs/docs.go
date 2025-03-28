// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://github.com/tule75/Booking-Go",
        "contact": {
            "name": "Tú Lê",
            "url": "http://github.com/tule75/Booking-Go",
            "email": "tulehd03@gmail.com"
        },
        "license": {
            "name": "Booking Go v1.0",
            "url": "http://github.com/tule75/Booking-Go"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/properties": {
            "post": {
                "description": "Create a new property",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties management"
                ],
                "summary": "Create a new property",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.PropertyCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/properties/filter": {
            "post": {
                "description": "Search Properties with filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties management"
                ],
                "summary": "Search Properties",
                "parameters": [
                    {
                        "description": "param",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.SearchPropertiesParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/properties/owner/{id}": {
            "get": {
                "description": "GetPropertiesByOwnerID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties management"
                ],
                "summary": "GetPropertiesByOwnerID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Owner ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit the number of rooms affected (default: 20)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset for starting position (default: 0)",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/properties/{id}": {
            "get": {
                "description": "GetPropertiesByID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties management"
                ],
                "summary": "GetPropertiesByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Owner ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            },
            "put": {
                "description": "Search Properties with filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties management"
                ],
                "summary": "Search Properties",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "param",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.PropertyUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Properties by its unique ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "properties management"
                ],
                "summary": "Delete Properties",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/rooms": {
            "post": {
                "description": "Create Room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room management"
                ],
                "summary": "Create Room",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.RoomCreateModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/rooms/property/{id}": {
            "get": {
                "description": "Get rooms from the system using property unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room management"
                ],
                "summary": "Get rooms by property ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Property ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit the number of rooms affected (default: 20)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset for starting position (default: 0)",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/rooms/{id}": {
            "get": {
                "description": "Get a room from the system using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room management"
                ],
                "summary": "Get a room by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            },
            "put": {
                "description": "Update rooms from the system using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room management"
                ],
                "summary": "Update room by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Property ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.RoomUpdateModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a room from the system using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room management"
                ],
                "summary": "Delete a room by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Room deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/users/create-password": {
            "post": {
                "description": "Verify OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account management"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.UserCreateRequestModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account management"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.LoginRequestModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register to get OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account management"
                ],
                "summary": "Register to get OTP",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.RegisterRequestModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        },
        "/users/verify-otp": {
            "post": {
                "description": "Verify OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account management"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestDTO.VerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.SearchPropertiesParams": {
            "type": "object",
            "properties": {
                "from_price": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "to_price": {
                    "type": "string"
                }
            }
        },
        "requestDTO.LoginRequestModel": {
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
        "requestDTO.PropertyCreateRequest": {
            "type": "object",
            "required": [
                "amenities",
                "location",
                "name",
                "price"
            ],
            "properties": {
                "amenities": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                }
            }
        },
        "requestDTO.PropertyUpdateRequest": {
            "type": "object",
            "required": [
                "amenities",
                "location",
                "name",
                "price"
            ],
            "properties": {
                "amenities": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                }
            }
        },
        "requestDTO.RegisterRequestModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "requestDTO.RoomCreateModel": {
            "type": "object",
            "properties": {
                "is_available": {
                    "type": "boolean"
                },
                "max_guests": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "property_id": {
                    "type": "string"
                }
            }
        },
        "requestDTO.RoomUpdateModel": {
            "type": "object",
            "properties": {
                "is_available": {
                    "type": "boolean"
                },
                "max_guests": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "property_id": {
                    "type": "string"
                }
            }
        },
        "requestDTO.UserCreateRequestModel": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/requestDTO.UserRole"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "requestDTO.UserRole": {
            "type": "string",
            "enum": [
                "CUSTOMER",
                "HOST"
            ],
            "x-enum-varnames": [
                "CUSTOMER",
                "HOST"
            ]
        },
        "requestDTO.VerifyRequest": {
            "type": "object",
            "properties": {
                "verify_code": {
                    "type": "string"
                },
                "verify_key": {
                    "type": "string"
                }
            }
        },
        "response.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/v1/2025",
	Schemes:          []string{},
	Title:            "Booking Go",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
