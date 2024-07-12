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
        "/sysOperationRecord/createSysOperationRecord": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysOperationRecord"
                ],
                "summary": "创建SysOperationRecord",
                "parameters": [
                    {
                        "description": "创建SysOperationRecord",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysOperationRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建SysOperationRecord",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
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
        "/sysOperationRecord/deleteSysOperationRecord": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysOperationRecord"
                ],
                "summary": "删除SysOperationRecord",
                "parameters": [
                    {
                        "description": "SysOperationRecord模型",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysOperationRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除SysOperationRecord",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
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
        "/sysOperationRecord/deleteSysOperationRecordByIds": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysOperationRecord"
                ],
                "summary": "批量删除SysOperationRecord",
                "parameters": [
                    {
                        "description": "批量删除SysOperationRecord",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.IdsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "批量删除SysOperationRecord",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
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
        "/sysOperationRecord/findSysOperationRecord": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysOperationRecord"
                ],
                "summary": "用id查询SysOperationRecord",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键ID",
                        "name": "ID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "代理",
                        "name": "agent",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求Body",
                        "name": "body",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "创建时间",
                        "name": "createdAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "错误信息",
                        "name": "error_message",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求ip",
                        "name": "ip",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "延迟",
                        "name": "latency",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求方法",
                        "name": "method",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求路径",
                        "name": "path",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "响应Body",
                        "name": "resp",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "请求状态",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "更新时间",
                        "name": "updatedAt",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用id查询SysOperationRecord",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        },
                                        "msg": {
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
        "/sysOperationRecord/getSysOperationRecordList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysOperationRecord"
                ],
                "summary": "分页获取SysOperationRecord列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键ID",
                        "name": "ID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "代理",
                        "name": "agent",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求Body",
                        "name": "body",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "创建时间",
                        "name": "createdAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "错误信息",
                        "name": "error_message",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求ip",
                        "name": "ip",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "延迟",
                        "name": "latency",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求方法",
                        "name": "method",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页大小",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求路径",
                        "name": "path",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "响应Body",
                        "name": "resp",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "请求状态",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "更新时间",
                        "name": "updatedAt",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.PageResult"
                                        },
                                        "msg": {
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
        "request.IdsReq": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {},
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
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
        "system.SysOperationRecord": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "agent": {
                    "description": "代理",
                    "type": "string"
                },
                "body": {
                    "description": "请求Body",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "error_message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "ip": {
                    "description": "请求ip",
                    "type": "string"
                },
                "latency": {
                    "description": "延迟",
                    "type": "string"
                },
                "method": {
                    "description": "请求方法",
                    "type": "string"
                },
                "path": {
                    "description": "请求路径",
                    "type": "string"
                },
                "resp": {
                    "description": "响应Body",
                    "type": "string"
                },
                "status": {
                    "description": "请求状态",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户id",
                    "type": "integer"
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
