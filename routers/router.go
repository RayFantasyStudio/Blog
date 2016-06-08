package routers

import (
	"github.com/RayFantasyStudio/blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/view", &controllers.ArticleViewController{})
	beego.Router("/article/create", &controllers.ArticleCreateController{})
	beego.Router("/article/modify", &controllers.ArticleModifyController{})

	beego.AutoRouter(&controllers.ArticleViewController{})

	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
}
