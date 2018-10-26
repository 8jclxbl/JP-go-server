package routers

import (
	"github.com/astaxie/beego"
	"JP-go-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/:action:string", &controllers.UserController{})
	beego.Router("/api/person/:action([\\w]+):person_id(&[\\w]+)?", &controllers.PersonController{})
}
