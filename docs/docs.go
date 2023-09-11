// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/restaurant/{id}": {
            "get": {
                "description": "Get reviews of a restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get reviews of a restaurant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Restaurant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Add rating to a restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add rating to a restaurant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Restaurant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Author Name",
                        "name": "author_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Rating",
                        "name": "rating",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Text",
                        "name": "text",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/restaurants": {
            "get": {
                "description": "Get all restaurants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all restaurants",
                "responses": {}
            },
            "post": {
                "description": "Create a new restaurant entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new restaurant",
                "parameters": [
                    {
                        "description": "Restaurant Name",
                        "name": "restaurantName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Photo",
                        "name": "photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Latitude",
                        "name": "lat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "number"
                        }
                    },
                    {
                        "description": "Longitude",
                        "name": "long",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "number"
                        }
                    },
                    {
                        "description": "Rating",
                        "name": "rating",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Restaurant created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Fields are empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to save restaurant",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/restaurants/{rating}": {
            "get": {
                "description": "Filter restaurants by rating",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Filter restaurants by rating",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rating",
                        "name": "rating",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Restaurant API",
	Description:      "This is a sample restaurant API with Swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
