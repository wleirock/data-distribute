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

	// 机构脚本管理
	beego.Router("/web/script/index", &controllers.ScriptController{}, "get:Index")
	beego.Router("/web/script/list", &controllers.ScriptController{}, "post:List")
	beego.Router("/web/script/add", &controllers.ScriptController{}, "get:Add")
	beego.Router("/web/script/save", &controllers.ScriptController{}, "post:Save")
	beego.Router("/web/script/delete", &controllers.ScriptController{}, "post:Delete")
	beego.Router("/web/script/hospitalList", &controllers.ScriptController{}, "get:GetHospitalList")
	beego.Router("/web/script/methodlList", &controllers.ScriptController{}, "get:GetMethodList")

	// 脚本文件
	beego.Router("/web/script/file/index", &controllers.ScriptFileController{}, "get:Index")
	beego.Router("/web/script/file/list", &controllers.ScriptFileController{}, "post:List")
	beego.Router("/web/script/file/upload", &controllers.ScriptFileController{}, "post:Upload")
	beego.Router("/web/script/file/delete", &controllers.ScriptFileController{}, "post:Delete")
	beego.Router("/web/script/file/download", &controllers.ScriptFileController{}, "get:Download")

	// 公共方法
	beego.Router("/web/method/index", &controllers.PublicMethodController{}, "get:Index")
	beego.Router("/web/method/list", &controllers.PublicMethodController{}, "post:List")
	beego.Router("/web/method/add", &controllers.PublicMethodController{}, "get:Add")
	beego.Router("/web/method/save", &controllers.PublicMethodController{}, "post:Save")
	beego.Router("/web/method/delete", &controllers.PublicMethodController{}, "post:Delete")

	// API
	beego.Router("/api/repportInfo/distribute", &controllers.ReportInfoController{}, "*:Distribute")

	// 注册过滤器
	beego.InsertFilter("/web/*", beego.BeforeRouter, LoginFilter)
}
