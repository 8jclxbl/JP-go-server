package routers

import (
	"github.com/astaxie/beego"
	"JP-go-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/reg/", &controllers.UserController{},"post:Reg")
	beego.Router("/api/login/", &controllers.UserController{},"post:Login")
	beego.Router("/api/logout/", &controllers.UserController{},"post:Logout")
	beego.Router("/api/update/", &controllers.UserController{},"post:Update")

}
