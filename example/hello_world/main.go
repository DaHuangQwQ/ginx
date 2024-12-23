package main

import (
	"github.com/DaHuangQwQ/ginx"
	"github.com/gin-gonic/gin"
)

type UserGetReq struct {
	ginx.Meta `method:"GET" path:"users/:id"`
	Id        int `json:"id" validate:"required,min=1,max=32"`
}

type UserGetRes struct {
	Code int `json:"code"`
}

func getUser(ctx *gin.Context, req UserGetReq) (ginx.Result[UserGetRes], error) {
	return ginx.Result[UserGetRes]{
		Code: 0,
		Msg:  "ok",
		Data: UserGetRes{
			Code: 1,
		},
	}, nil
}

func main() {
	server := ginx.NewServer(":8080")
	server.Handle(ginx.Wrap[UserGetReq, UserGetRes](getUser))
	_ = server.Start()
}
