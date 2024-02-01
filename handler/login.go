package handler

import (
	"context"
	"gf_gtoken/model"
)

type ListHandler struct{}

/*
Sign 此函数的请求/响应格式为固定写法 否则注册路由时注册不上并且响应时无法正确响应如下结构:

	{
		code:0,
		message:"提示信息",
		data:{}
	}
*/
func (h *ListHandler) Sign(ctx context.Context, req *model.SignReq) (res *model.SignRes, err error) {
	return nil, nil
}
