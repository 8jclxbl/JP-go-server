package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/reg/", &controllers.RegController{})
	beego.Router("/api/login/", &controllers.LoginController{})
	beego.Router("/api/logout/", &controllers.LogoutController{})
	beego.Router("/api/update/", &controllers.UpdateController{})
}
