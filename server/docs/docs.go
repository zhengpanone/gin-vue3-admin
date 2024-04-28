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
            "name": "go-swagger帮助文档",
            "url": "https://github.com/swaggo/swag/blob/master/README_zh-CN.md"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/api/admin/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "operationId": "/v1/api/admin/login",
                "parameters": [
                    {
                        "description": "用户名,密码,验证码,验证码ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginParam"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/admin/logout": {
            "post": {
                "description": "退出登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户退出",
                "operationId": "/v1/api/admin/logout",
                "parameters": [
                    {
                        "description": "用户名,密码,验证码,验证码ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginParam"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/admin/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户注册",
                "operationId": "/v1/api/admin/register",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterParam"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/captcha": {
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
                    "Base"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "生成验证码,返回包括随机数id,base64,验证码长度",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.SysCaptchaResponse"
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
        "/v1/api/changePassword": {
            "post": {
                "description": "更改密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "更改密码",
                "operationId": "/v1/user/changePassword",
                "responses": {}
            }
        },
        "/v1/api/config": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "获取配置",
                "responses": {}
            }
        },
        "/v1/api/jwt/jsonInBlacklist": {
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
                    "Jwt"
                ],
                "summary": "jwt加入黑名单",
                "responses": {
                    "200": {
                        "description": "jwt加入黑名单",
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
        "/v1/api/menu/addBaseMenu": {
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
                    "Menu"
                ],
                "summary": "新增菜单",
                "parameters": [
                    {
                        "description": "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysBaseMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "新增菜单",
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
        "/v1/api/menu/getMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据用户ID获取菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AuthorityMenu"
                ],
                "summary": "获取用户动态路由",
                "parameters": [
                    {
                        "description": "空",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Empty"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取用户动态路由,返回包括系统菜单详情列表",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.SysMenusResponse"
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
        "/v1/api/menu/getMenuList": {
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
                    "Menu"
                ],
                "summary": "分页获取基础menu列表",
                "parameters": [
                    {
                        "description": "页码, 每页大小",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分页获取基础menu列表,返回包括列表,总数,页码,每页数量",
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
        },
        "/v1/api/role/addRole": {
            "post": {
                "description": "创建角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RoleApi"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RoleParam"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/role/delete/:id": {
            "delete": {
                "description": "删除角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RoleApi"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "角色ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/v1/api/role/pageRole": {
            "post": {
                "description": "分页获取角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RoleApi"
                ],
                "summary": "分页获取角色",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/role/update": {
            "put": {
                "description": "update role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RoleApi"
                ],
                "summary": "更新角色",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "request.Empty": {
            "type": "object"
        },
        "request.LoginParam": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "captcha": {
                    "description": "验证码",
                    "type": "string"
                },
                "captchaID": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.PageInfo": {
            "type": "object",
            "properties": {
                "keyword": {
                    "description": "关键字",
                    "type": "string"
                },
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页大小",
                    "type": "integer"
                }
            }
        },
        "request.RegisterParam": {
            "type": "object",
            "required": [
                "nickName",
                "password",
                "phone",
                "username"
            ],
            "properties": {
                "address": {
                    "description": "地址",
                    "type": "string"
                },
                "birthday": {
                    "description": "生日",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 3
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号码",
                    "type": "string"
                },
                "roleID": {
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.RoleParam": {
            "type": "object",
            "required": [
                "roleName"
            ],
            "properties": {
                "parentId": {
                    "description": "父角色ID",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名",
                    "type": "string"
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
                    "description": "业务状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "响应数据"
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "response.SysCaptchaResponse": {
            "type": "object",
            "properties": {
                "captcha_id": {
                    "type": "string"
                },
                "captcha_length": {
                    "type": "integer"
                },
                "pic_path": {
                    "type": "string"
                }
            }
        },
        "response.SysMenusResponse": {
            "type": "object",
            "properties": {
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysMenu"
                    }
                }
            }
        },
        "system.Meta": {
            "type": "object",
            "properties": {
                "activeName": {
                    "type": "string"
                },
                "closeTab": {
                    "description": "自动关闭tab",
                    "type": "boolean"
                },
                "defaultMenu": {
                    "description": "是否是基础路由（开发中）",
                    "type": "boolean"
                },
                "icon": {
                    "description": "菜单图标",
                    "type": "string"
                },
                "keepAlive": {
                    "description": "是否缓存",
                    "type": "boolean"
                },
                "title": {
                    "description": "菜单名",
                    "type": "string"
                }
            }
        },
        "system.SysAuthority": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "UpdateAt": {
                    "description": "使用时间戳毫秒数填充更新时间",
                    "type": "string"
                },
                "authorityId": {
                    "description": "角色ID",
                    "type": "integer"
                },
                "authorityName": {
                    "description": "角色名",
                    "type": "string"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "createAt": {
                    "description": "使用时间戳秒数填充创建时间",
                    "type": "string"
                },
                "dataAuthorityId": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "defaultRouter": {
                    "description": "默认菜单(默认dashboard)",
                    "type": "string"
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenu"
                    }
                },
                "parentId": {
                    "description": "父角色ID",
                    "type": "integer"
                }
            }
        },
        "system.SysBaseMenu": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "UpdateAt": {
                    "description": "使用时间戳毫秒数填充更新时间",
                    "type": "string"
                },
                "authoritys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenu"
                    }
                },
                "component": {
                    "type": "string"
                },
                "createAt": {
                    "description": "使用时间戳秒数填充创建时间",
                    "type": "string"
                },
                "hidden": {
                    "type": "boolean"
                },
                "menuBtn": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuBtn"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/system.Meta"
                },
                "name": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuParameter"
                    }
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "system.SysBaseMenuBtn": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "UpdateAt": {
                    "description": "使用时间戳毫秒数填充更新时间",
                    "type": "string"
                },
                "createAt": {
                    "description": "使用时间戳秒数填充创建时间",
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sysBaseMenuID": {
                    "type": "integer"
                }
            }
        },
        "system.SysBaseMenuParameter": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "UpdateAt": {
                    "description": "使用时间戳毫秒数填充更新时间",
                    "type": "string"
                },
                "createAt": {
                    "description": "使用时间戳秒数填充创建时间",
                    "type": "string"
                },
                "key": {
                    "description": "地址栏携带参数的key",
                    "type": "string"
                },
                "sysBaseMenuID": {
                    "type": "integer"
                },
                "type": {
                    "description": "地址栏携带参数为params还是query",
                    "type": "string"
                },
                "value": {
                    "description": "地址栏携带参数的值",
                    "type": "string"
                }
            }
        },
        "system.SysMenu": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "UpdateAt": {
                    "description": "使用时间戳毫秒数填充更新时间",
                    "type": "string"
                },
                "authoritys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "btns": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysMenu"
                    }
                },
                "component": {
                    "type": "string"
                },
                "createAt": {
                    "description": "使用时间戳秒数填充创建时间",
                    "type": "string"
                },
                "hidden": {
                    "type": "boolean"
                },
                "menuBtn": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuBtn"
                    }
                },
                "menu_id": {
                    "type": "integer"
                },
                "meta": {
                    "$ref": "#/definitions/system.Meta"
                },
                "name": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuParameter"
                    }
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "perm": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysBaseMenuParameter"
                    }
                },
                "sort": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gin api learn",
	Description:      "This is gin learn api docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
