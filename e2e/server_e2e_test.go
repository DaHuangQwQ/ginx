//go:build e2e

package e2e

import (
	"encoding/json"
	"github.com/DaHuangQwQ/ginx"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer(t *testing.T) {
	server := ginx.NewServer()
	server.Handle(ginx.Wrap[userGetReq, userGetRes](getUser))

	marshal, err := json.Marshal(ginx.Oai.Description())
	if err != nil {
		return
	}
	println(string(marshal))
}

func getUser(ctx *gin.Context, req userGetReq) (ginx.Result[userGetRes], error) {
	return ginx.Result[userGetRes]{
		Code: 0,
		Msg:  "ok",
		Data: userGetRes{},
	}, nil
}

type userGetReq struct {
	ginx.Meta `method:"GET" path:"users/:id"`
	Id        int `json:"id" validate:"required,min=1,max=32"`
}

type userGetRes struct {
	ginx.Meta
}
