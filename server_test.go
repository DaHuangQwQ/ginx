package ginx

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gin-gonic/gin"
)

func TestServer_marshalSpec(t *testing.T) {
	testCases := []struct {
		name string

		server *Server

		wantRes string
		wantErr error
	}{
		{
			name: "normal",
			server: func() *Server {
				server := NewServer(":8081")
				server.Handle(Wrap[UserGetReq, UserGetRes](getUser))
				return server
			}(),

			wantRes: ``,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.server.MarshalSpec()
			require.NoError(t, err)
			println(string(b))
			require.Equal(t, string(b), `{
	"components": {
		"schemas": {
			"UserGetReq": {
				"description": "UserGetReq schema",
				"properties": {
					"id": {
						"maximum": 32,
						"minimum": 1,
						"type": "integer"
					}
				},
				"required": [
					"id"
				],
				"type": "object"
			},
			"UserGetRes": {
				"description": "UserGetRes schema",
				"properties": {
					"code": {
						"type": "integer"
					}
				},
				"type": "object"
			}
		}
	},
	"info": {
		"description": "123",
		"title": "OpenAPI",
		"version": "0.0.1"
	},
	"openapi": "3.0.1",
	"paths": {
		"/users/:id": {
			"get": {
				"description": "#### Controller: \n\n`+"`/users/:id`"+`\n\n---\n\n",
				"operationId": "GET_/users/:id",
				"requestBody": {
					"content": {
						"*/*": {
							"schema": {
								"$ref": "#/components/schemas/UserGetReq"
							}
						}
					},
					"description": "Request body for ginx.UserGetReq",
					"required": true
				},
				"responses": {
					"200": {
						"content": {
							"application/json": {
								"schema": {
									"$ref": "#/components/schemas/UserGetRes"
								}
							},
							"application/xml": {
								"schema": {
									"$ref": "#/components/schemas/UserGetRes"
								}
							}
						},
						"description": "OK"
					},
					"default": {
						"description": ""
					}
				},
				"summary": "/users/:id"
			}
		}
	}
}`)
		})
	}
}

type UserGetReq struct {
	Meta `method:"GET" path:"/users/:id"`
	Id   int `json:"id" validate:"required,min=1,max=32"`
}

type UserGetRes struct {
	Code int `json:"code"`
}

func getUser(ctx *gin.Context, req UserGetReq) (Result[UserGetRes], error) {
	return Result[UserGetRes]{
		Code: 0,
		Msg:  "ok",
		Data: UserGetRes{
			Code: 1,
		},
	}, nil
}

func TestServer_RegisterOpenAPIRoutes(t *testing.T) {
	//server := NewServer(":8081")
	//server.Handle(Wrap[UserGetReq, UserGetRes](getUser))
	//server.RegisterOpenAPIRoutes("/openapi")
	//err := server.Start()
	//if err != nil {
	//	return
	//}
}
