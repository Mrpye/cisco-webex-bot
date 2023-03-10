// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://github.com/Mrpye/cisco-webex-bot"
        },
        "license": {
            "name": "Apache 2.0 licensed",
            "url": "https://github.com/Mrpye/cisco-webex-bot/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Check API Endpoint",
                "operationId": "check-api-endpoint",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/send_message": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Send  a message to Webex text format",
                "operationId": "post-send-message-text",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/body_types.SendMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message sent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/send_message_md": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Send  a message to Webex markdown format",
                "operationId": "post-send-message-markdown",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/body_types.SendMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message sent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "body_types.SendMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "cisco-webex-bot",
	Description:      "cisco-webex-bot is a CLI application written in Golang that gives the ability to send notifications to a Webex room that the bot is subscribed to via a rest api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
