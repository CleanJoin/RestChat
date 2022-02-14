// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/CleanJoin/RestChat/",
        "contact": {
            "name": "Github.com",
            "url": "https://github.com/CleanJoin/RestChat/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/health": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "Show the status of server.",
                "responses": {}
            }
        },
        "/api/login": {
            "post": {
                "description": "Вход пользователя в чат",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "loginHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restchat.RequestUser"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/logout": {
            "post": {
                "description": "Выход пользователя из чата",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "logoutHandler",
                "parameters": [
                    {
                        "description": "User ApiToken",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restchat.RequestApiToken"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/members": {
            "post": {
                "description": "Получить список online пользователей",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "membersHandler",
                "parameters": [
                    {
                        "description": "User ApiToken",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restchat.RequestApiToken"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/message": {
            "post": {
                "description": "Отправить сообщение",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "messageHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "MessageSend",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restchat.RequestMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/messages": {
            "post": {
                "description": "Получить список сообщений",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "messagesHandler",
                "parameters": [
                    {
                        "description": "User ApiToken",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restchat.RequestApiToken"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/user": {
            "post": {
                "description": "Регистарция пользователя",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "userHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restchat.RequestUser"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "restchat.RequestApiToken": {
            "type": "object",
            "properties": {
                "api_token": {
                    "type": "string"
                }
            }
        },
        "restchat.RequestMessage": {
            "type": "object",
            "properties": {
                "api_token": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "restchat.RequestUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger RestChat",
	Description:      "This is a sample server Rest API Server Chat.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
