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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/GetTicketInfoA": {
            "post": {
                "description": "Respond with Ticket info.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "GetTicketInfo input body",
                        "name": "GetTicketInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTicketInfo_input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTicketInfo"
                        }
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "401": {
                        "description": "token in not valid or role is not admin"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/GetTicketInfoU": {
            "get": {
                "description": "Respond with User ticket info.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "GetTicketInfo input body",
                        "name": "GetTicketInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTicketInfo_input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTicketInfo"
                        }
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "401": {
                        "description": "Token is not valid"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/getTicketsListA": {
            "get": {
                "description": "Respond with admins tickets list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTicketsList_A"
                        }
                    },
                    "401": {
                        "description": "token in not valid or role is not admin"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/getTicketsListU": {
            "get": {
                "description": "Respond with users tickets list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTicketsList_U"
                        }
                    },
                    "401": {
                        "description": "Token is not valid"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/getcaptcha": {
            "post": {
                "description": "Get captcha PNG from server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get random captcha id",
                        "name": "refId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Captcha image generated"
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Respond with Token string as Json if login was succesfull.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Login input body",
                        "name": "LoginInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.login_input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successfull",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "401": {
                        "description": "username and password or captcha is not match"
                    },
                    "408": {
                        "description": "Captcha expired"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/messageForTicketU": {
            "post": {
                "description": "Respond with Json body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "MessageForTicketing input body",
                        "name": "MessageForTicketing",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageForTicketing_input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Succesfully message saved"
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "401": {
                        "description": "token in not valid or role is not admin"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "user registration.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Register input body",
                        "name": "RegisterInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Register_Input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "0: This Id is already registered , 1: This Username is already registered , 2: Register was successfull"
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "408": {
                        "description": "Captcha expired"
                    }
                }
            }
        },
        "/api/resetpassword": {
            "post": {
                "description": "Changes the user's password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Reset Password Body",
                        "name": "ResetPasswordInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Reset_password"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "password was successfully changed"
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "401": {
                        "description": "Invalid password for user"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/ticketing": {
            "post": {
                "description": "Get user's Message.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "Message was successfully received"
                    },
                    "400": {
                        "description": "Empty Body"
                    },
                    "401": {
                        "description": "Null token"
                    },
                    "500": {
                        "description": "Problem in saving message"
                    }
                }
            }
        },
        "/api/tokenValidation": {
            "post": {
                "description": "Respond with user information such as username, SSN, Role, Email and phone.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get token string",
                        "name": "BearerToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token is valid"
                    },
                    "401": {
                        "description": "Invalid token"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.GetTicketInfo": {
            "type": "object",
            "properties": {
                "textTime": {
                    "type": "string"
                },
                "textsForTicket": {
                    "type": "string"
                },
                "ticketMessage": {
                    "type": "string"
                },
                "ticketSubject": {
                    "type": "string"
                }
            }
        },
        "controllers.GetTicketInfo_input": {
            "type": "object",
            "required": [
                "ticketid"
            ],
            "properties": {
                "ticketid": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "controllers.GetTicketsList_A": {
            "type": "object",
            "properties": {
                "hasbeenRead": {
                    "type": "integer"
                },
                "ticketId": {
                    "type": "integer"
                },
                "ticketsMessage": {
                    "type": "string"
                },
                "ticketsSubject": {
                    "type": "string"
                }
            }
        },
        "controllers.GetTicketsList_U": {
            "type": "object",
            "properties": {
                "hasbeenRead": {
                    "type": "integer"
                },
                "ticketID": {
                    "type": "integer"
                },
                "ticketMessage": {
                    "type": "string"
                },
                "ticketSubject": {
                    "type": "string"
                }
            }
        },
        "controllers.MessageForTicketing_input": {
            "type": "object",
            "required": [
                "text",
                "ticketid"
            ],
            "properties": {
                "text": {
                    "type": "string"
                },
                "ticketid": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "controllers.Register_Input": {
            "type": "object",
            "required": [
                "captchaCode",
                "captchaId",
                "email",
                "id",
                "password",
                "phone",
                "username"
            ],
            "properties": {
                "captchaCode": {
                    "type": "string"
                },
                "captchaId": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "example": "0"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.Reset_password": {
            "type": "object",
            "required": [
                "newPassword",
                "oldPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "controllers.login_input": {
            "type": "object",
            "required": [
                "captchaCode",
                "captchaId",
                "pwd",
                "username"
            ],
            "properties": {
                "captchaCode": {
                    "type": "string"
                },
                "captchaId": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
