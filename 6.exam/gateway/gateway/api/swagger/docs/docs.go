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
        "/api/notifications/unread": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This method retrieves all unread notifications for the user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NOTIFICATIONS"
                ],
                "summary": "Get unread notifications",
                "responses": {
                    "200": {
                        "description": "List of unread notifications",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/notification.NotifyList"
                            }
                        }
                    },
                    "403": {
                        "description": "Permission Denied",
                        "schema": {}
                    },
                    "500": {
                        "description": "Unable to get unread notifications",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "notification.Notify": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "report": {
                    "$ref": "#/definitions/notification.Report"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "notification.NotifyList": {
            "type": "object",
            "properties": {
                "notifyList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/notification.Notify"
                    }
                }
            }
        },
        "notification.Report": {
            "type": "object",
            "properties": {
                "expenses": {
                    "type": "number"
                },
                "income": {
                    "type": "number"
                },
                "netSavings": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "PERSONAL FINANCE MANAGEMENT",
	Description:      "This swagger UI was created to manage personal finance",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
