package token

import (
	"gf_gtoken/consts"
	"gf_gtoken/model"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

type Auth struct{}

func NewAuthToken() *Auth {
	return &Auth{}
}

var BfToken gtoken.GfToken
var ctx = gctx.New()

func (c *Auth) Login() {
	newAuth := new(Auth)
	// BfToken = new(gtoken.GfToken)
	BfToken = gtoken.GfToken{
		ServerName:      "_gf_web",
		Timeout:         g.Cfg().MustGet(ctx, "gToken.Timeout").Int(),    // 超时时间
		CacheMode:       g.Cfg().MustGet(ctx, "gToken.CacheMode").Int8(), // 缓存模式
		LoginPath:       "/login",                                        // 登录路径
		LogoutPath:      "/logout",                                       // 退出
		LoginBeforeFunc: newAuth.loginFunc,                               // 登录校验的函数
		LoginAfterFunc:  newAuth.loginAfterFunc,                          // 登录完成做哪些操作
		AuthPaths:       g.SliceStr{"/login"},                            // 这里是按照前缀拦截，拦截/user /user/list /user/add ...
		// AuthExcludePaths: g.SliceStr{"/user/info", "/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		GlobalMiddleware: true, // 开启全局拦截
	}
	if err := BfToken.Start(); err != nil {
		g.Log().Error(ctx, err)
		panic(err)
	}
}

// 下面的报错是因为 sql 与结构体报错 换成自己的就可
func (c *Auth) loginFunc(r *ghttp.Request) (username string, data interface{}) {
	req := new(model.UserSignInReq)
	if err := r.Parse(req); err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.UsernameOrPasswordNotNull))
		r.ExitAll()
		return
	}
	var count int
	newData := new(entity.Users)
	err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, req.Username).ScanAndCount(&newData, &count, true)
	if count == 0 {
		r.Response.WriteJson(gtoken.Fail(consts.UserNotExist))
		r.ExitAll()
		return
	}
	if err != nil {
		g.Log().Error(gctx.New(), err)
		return "", err
	}
	if EncryptPassword(req.Password) != newData.Password || req.Username != newData.Username {
		r.Response.WriteJson(gtoken.Fail(consts.UsernameOrPasswordErr))
		r.ExitAll()
		return
	}
	// name 作为 KEY 键值返回,data 登录成功数据结构体
	return gconv.String(newData.Id), newData
}

func (c *Auth) loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
		return
	} else {
		data := new(entity.Users)
		id := respData.Get(consts.UserKey).Int64()
		err := dao.Users.Ctx(ctx).WherePri(id).Scan(&data)
		if err != nil {
			return
		}
		res := &model.UserSignInRes{
			Token:      respData.GetString(consts.Token),
			Id:         data.Id,
			Username:   data.Username,
			Nickname:   data.Nickname,
			CreateTime: data.CreateTime,
			UpdateTime: data.UpdateTime,
			DeleteTime: data.DeleteTime,
		}
		respData.Data = res
		r.Response.WriteJson(respData)
		return
	}
}

// EncryptPassword 密码加密
func EncryptPassword(password string) (newPsw string) {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password))
}
