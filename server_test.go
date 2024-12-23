package ginx

import (
	"testing"

	"github.com/DaHuangQwQ/ginx/middleware/prometheus"

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
		})
	}
}

type UserGetReq struct {
	Meta `method:"GET" path:"/users/:id"`
	Id   int  `json:"id" validate:"required,min=1,max=32"`
	Demo Demo `json:"demo"`
}

type UserGetRes struct {
	Code int `json:"code"`
}

type Demo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
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

func TestServer_RegisterOpenAPIMiddleware(t *testing.T) {
	server := NewServer(":8081")
	builder := prometheus.Builder{
		Namespace:  "test",
		Subsystem:  "test",
		Name:       "user",
		InstanceId: "1",
		Help:       "1",
	}
	server.Use(builder.BuildActiveRequest())
	server.Use(builder.BuildResponseTime())
	server.RegisterOpenAPIMiddleware()
	server.Handle(Wrap[UserGetReq, UserGetRes](getUser))

	server.RegisterOpenAPIRoutes("/openapi")
	//err := server.Start()
	//if err != nil {
	//	return
	//}
}
