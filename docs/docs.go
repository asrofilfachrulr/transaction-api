// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Developer",
            "email": "riidloa@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customer": {
            "post": {
                "description": "add new customer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "add new customer",
                "parameters": [
                    {
                        "description": "Register new customer",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.RegisterCustomerInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/web.NormalResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/web.RegisterCustomerOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/customers": {
            "get": {
                "description": "get all customers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/web.NormalResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/web.CustomerOutput"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/order": {
            "post": {
                "description": "add new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "add new user",
                "parameters": [
                    {
                        "description": "Entry new order",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.PostOrderInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/web.NormalResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/web.PostOrderOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "web.CustomerOutput": {
            "type": "object",
            "properties": {
                "customer_address_id": {
                    "type": "integer"
                },
                "customer_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "web.NormalResp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "any"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "web.PostOrderInput": {
            "type": "object"
        },
        "web.PostOrderOutput": {
            "type": "object"
        },
        "web.RegisterCustomerInput": {
            "type": "object",
            "required": [
                "address",
                "name"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "web.RegisterCustomerOutput": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "customer_address_id": {
                    "type": "integer"
                },
                "customer_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Transaction API",
	Description:      "API provide simple transaction",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
