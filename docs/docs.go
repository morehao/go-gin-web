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
        "/app/user/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "创建用户",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtoUser.UserCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0,\"data\": \"ok\",\"msg\": \"success\"}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.DefaultRender"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtoUser.UserCreateResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/user/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "删除用户",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtoUser.UserDeleteReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0,\"data\": \"ok\",\"msg\": \"删除成功\"}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.DefaultRender"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/user/detail": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据自增id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0,\"data\": \"ok\",\"msg\": \"success\"}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.DefaultRender"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtoUser.UserDetailResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/user/pageList": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户列表分页",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 1000,
                        "type": "integer",
                        "description": "每页数据条数",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0,\"data\": \"ok\",\"msg\": \"success\"}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.DefaultRender"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtoUser.UserPageListResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/user/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "description": "修改用户",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtoUser.UserUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0,\"data\": \"ok\",\"msg\": \"修改成功\"}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.DefaultRender"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
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
        "dto.DefaultRender": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "dtoUser.UserCreateReq": {
            "type": "object",
            "properties": {
                "companyId": {
                    "description": "公司id",
                    "type": "integer"
                },
                "departmentId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "name": {
                    "description": "姓名",
                    "type": "string"
                }
            }
        },
        "dtoUser.UserCreateResp": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "数据自增id",
                    "type": "integer"
                }
            }
        },
        "dtoUser.UserDeleteReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "description": "数据自增id",
                    "type": "integer"
                }
            }
        },
        "dtoUser.UserDetailResp": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "companyId": {
                    "description": "公司id",
                    "type": "integer"
                },
                "createdBy": {
                    "description": "创建人id",
                    "type": "integer"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "departmentId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "id": {
                    "description": "数据自增id",
                    "type": "integer"
                },
                "name": {
                    "description": "姓名",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "更新人id",
                    "type": "integer"
                },
                "updatedTime": {
                    "description": "更新时间",
                    "type": "integer"
                }
            }
        },
        "dtoUser.UserPageListItem": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "companyId": {
                    "description": "公司id",
                    "type": "integer"
                },
                "createdBy": {
                    "description": "创建人id",
                    "type": "integer"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "departmentId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "id": {
                    "description": "数据自增id",
                    "type": "integer"
                },
                "name": {
                    "description": "姓名",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "更新人id",
                    "type": "integer"
                },
                "updatedTime": {
                    "description": "更新时间",
                    "type": "integer"
                }
            }
        },
        "dtoUser.UserPageListResp": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "数据列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtoUser.UserPageListItem"
                    }
                },
                "total": {
                    "description": "数据总条数",
                    "type": "integer"
                }
            }
        },
        "dtoUser.UserUpdateReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "companyId": {
                    "description": "公司id",
                    "type": "integer"
                },
                "departmentId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "id": {
                    "description": "数据自增id",
                    "type": "integer"
                },
                "name": {
                    "description": "姓名",
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
