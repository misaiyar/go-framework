package routers

import (
	"project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.MainController{},"post:Login")
    beego.AutoRouter(&controllers.AdminController{})
    beego.AutoRouter(&controllers.SystemController{})
    beego.AutoRouter(&controllers.SysuserController{})
    beego.AutoRouter(&controllers.ModuleController{})
    beego.AutoRouter(&controllers.RoleController{})
}
