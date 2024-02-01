package cmd

import (
	"context"
	"gf_gtoken/handler"
	"gf_gtoken/token"
	"github.com/gogf/gf/v2/frame/g"
)

func MainRun() {
	s := g.Server()
	var ctx context.Context
	v1 := s.Group("/api/v1")
	// 这里先调用函数(初始化/配置信息等一系列化的 token 操作)使 token 生成 并响应给前端
	token.NewAuthToken().Login()
	// 假如想要注册登录路由 这里不用像注册列表路由那样 因为登录的逻辑在 token 文件 所以这里直接把 token 注册为中间件 登录接口会自动注册为路由
	err := token.BfToken.Middleware(ctx, v1)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	v1.POST("/list", handler.ListHandler{}) // 注册列表路由
}
