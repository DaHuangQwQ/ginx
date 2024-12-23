# ginx

```shell
go get github.com/DaHuangQwQ/ginx
```
样例代码
https://github.com/DaHuangQwQ/ginx/tree/main/example
1. 基于反射实现的自动 api 文档生成
```go
type UserGetReq struct {
	ginx.Meta `method:"GET" path:"users/:id"`
	Id        int `json:"id" validate:"required,min=1,max=32"`
}
```
2. jwt中间件
3. 限流中间件
4. 可观测中间件
5. 简化代码

```go
package main

import (
	"github.com/DaHuangQwQ/ginx"
	"github.com/gin-gonic/gin"
)

type UserGetReq struct {
	ginx.Meta `method:"GET" path:"/users/:id"`
	Id        int `json:"id" uri:"id" validate:"required,min=1,max=32"`
}

type UserGetRes struct {
	Code int `json:"code"`
}

func getUser(ctx *gin.Context, req UserGetReq) (ginx.Result[UserGetRes], error) {
	return ginx.Result[UserGetRes]{
		Code: 0,
		Msg:  "ok",
		Data: UserGetRes{
			Code: req.Id,
		},
	}, nil
}

func main() {
	server := ginx.NewServer(":8080")
	server.Handle(ginx.Wrap[UserGetReq, UserGetRes](getUser))
	_ = server.Start()
}

```