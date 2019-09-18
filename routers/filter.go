package routers

import "github.com/astaxie/beego/context"

//LoginFilter 登录验证过滤器
var LoginFilter = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("username").(string)
	if !ok {
		ctx.Redirect(302, "/")
	}
}
