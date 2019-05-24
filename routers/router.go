package routers

import (
	"wleirock/data-distribute/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/repportInfo/distribute", &controllers.ReportInfoController{}, "*:Distribute")
}
