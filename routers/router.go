package routers

import (
	"github.com/RayFantasyStudio/blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
}
