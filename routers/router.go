package routers

import (
	"github.com/astaxie/beego"
	"JP-go-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/:action:string", &controllers.UserController{})
	beego.Router("/api/person/:action([\\w]+):person_id(&[\\w]+)", &controllers.PersonController{},"get:Get")
	beego.Router("/api/person/:action([\\w]+)", &controllers.PersonController{},"post:Post")
	beego.Router("/api/event/:action([\\w]+):event_id(&[\\w]+)", &controllers.EventController{},"get:Get")
	beego.Router("/api/event/:action",&controllers.EventController{},"post:Post")
	beego.Router("/api/file/upload",&controllers.FileController{},"get:Get;post:Upload")
	beego.Router("/api/file/delete&:file_url*",&controllers.FileController{},"get:Delete")
}
