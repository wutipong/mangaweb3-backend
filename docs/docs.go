// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/browse": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/browse.browseRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/browse.browseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/browse/rescan_library": {
            "get": {
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/browse.recreateThumbnailsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/browse/thumbnail": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the item",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/tag/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tag.listRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tag.listResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/tag/set_favorite": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tag.setFavoriteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tag.setFavoriteResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/tag/thumbnail": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "tag",
                        "name": "tag",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/view": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.viewRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.viewResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/view/download": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the file",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/view/get_image": {
            "get": {
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the item",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "width",
                        "name": "width",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "height",
                        "name": "height",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "index",
                        "name": "i",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/view/set_favorite": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.setFavoriteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.setFavoriteResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/view/update_cover": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.updateCoverRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.updateCoverResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "browse.browseRequest": {
            "type": "object",
            "properties": {
                "favorite_only": {
                    "type": "boolean"
                },
                "item_per_page": {
                    "type": "integer"
                },
                "order": {
                    "$ref": "#/definitions/meta.SortOrder"
                },
                "page": {
                    "type": "integer"
                },
                "search": {
                    "type": "string"
                },
                "sort": {
                    "$ref": "#/definitions/meta.SortField"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "browse.browseResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Meta"
                    }
                },
                "request": {
                    "$ref": "#/definitions/browse.browseRequest"
                },
                "tag_favorite": {
                    "type": "boolean"
                },
                "total_page": {
                    "type": "integer"
                }
            }
        },
        "browse.recreateThumbnailsResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "boolean"
                }
            }
        },
        "browse.rescanLibraryResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "boolean"
                }
            }
        },
        "ent.Meta": {
            "type": "object",
            "properties": {
                "create_time": {
                    "description": "CreateTime holds the value of the \"create_time\" field.",
                    "type": "string"
                },
                "favorite": {
                    "description": "Favorite holds the value of the \"favorite\" field.",
                    "type": "boolean"
                },
                "file_indices": {
                    "description": "FileIndices holds the value of the \"file_indices\" field.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "read": {
                    "description": "Read holds the value of the \"read\" field.",
                    "type": "boolean"
                },
                "tags": {
                    "description": "Tags holds the value of the \"tags\" field.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "ent.Tag": {
            "type": "object",
            "properties": {
                "favorite": {
                    "description": "Favorite holds the value of the \"favorite\" field.",
                    "type": "boolean"
                },
                "hidden": {
                    "description": "Hidden holds the value of the \"hidden\" field.",
                    "type": "boolean"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "errors.Error": {
            "type": "object",
            "properties": {
                "cause": {},
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "message": {
                    "type": "string",
                    "example": "unknown error."
                }
            }
        },
        "meta.SortField": {
            "type": "string",
            "enum": [
                "name",
                "createTime"
            ],
            "x-enum-varnames": [
                "SortFieldName",
                "SortFieldCreateTime"
            ]
        },
        "meta.SortOrder": {
            "type": "string",
            "enum": [
                "ascending",
                "descending"
            ],
            "x-enum-varnames": [
                "SortOrderAscending",
                "SortOrderDescending"
            ]
        },
        "tag.listRequest": {
            "type": "object",
            "properties": {
                "favorite_only": {
                    "type": "boolean"
                },
                "item_per_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "tag.listResponse": {
            "type": "object",
            "properties": {
                "request": {
                    "$ref": "#/definitions/tag.listRequest"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Tag"
                    }
                },
                "total_page": {
                    "type": "integer"
                }
            }
        },
        "tag.setFavoriteRequest": {
            "type": "object",
            "properties": {
                "favorite": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "tag.setFavoriteResponse": {
            "type": "object",
            "properties": {
                "favorite": {
                    "type": "boolean"
                },
                "request": {
                    "$ref": "#/definitions/tag.setFavoriteRequest"
                }
            }
        },
        "view.setFavoriteRequest": {
            "type": "object",
            "properties": {
                "favorite": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "view.setFavoriteResponse": {
            "type": "object",
            "properties": {
                "favorite": {
                    "type": "boolean"
                },
                "request": {
                    "$ref": "#/definitions/view.setFavoriteRequest"
                }
            }
        },
        "view.updateCoverRequest": {
            "type": "object",
            "properties": {
                "index": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "view.updateCoverResponse": {
            "type": "object",
            "properties": {
                "request": {
                    "$ref": "#/definitions/view.updateCoverRequest"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "view.viewRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "view.viewResponse": {
            "type": "object",
            "properties": {
                "favorite": {
                    "type": "boolean"
                },
                "indices": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "type": "string"
                },
                "request": {
                    "$ref": "#/definitions/view.viewRequest"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "version": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Mangaweb3 API",
	Description:      "API Server for Mangaweb",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
