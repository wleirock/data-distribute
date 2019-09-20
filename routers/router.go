package routers

import (
	"wleirock/data-distribute/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 账号
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/account/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/account/logout", &controllers.AccountController{}, "get:Logout")

	// 脚本管理
	beego.Router("/web/script/index", &controllers.ScriptController{}, "get:Index")
	beego.Router("/web/script/list", &controllers.ScriptController{}, "post:List")
	beego.Router("/web/script/add", &controllers.ScriptController{}, "get:Add")
	beego.Router("/web/script/save", &controllers.ScriptController{}, "post:Save")

	// API
	beego.Router("/api/repportInfo/distribute", &controllers.ReportInfoController{}, "*:Distribute")

	// 注册过滤器
	beego.InsertFilter("/web/*", beego.BeforeRouter, LoginFilter)
}
