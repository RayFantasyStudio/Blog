package main

import (
	_ "github.com/RayFantasyStudio/blog/routers"
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/utils"
	"github.com/RayFantasyStudio/blog/models"
)

func main() {
	beego.AddFuncMap("SinceTime", utils.SinceTime)
	ar,_,err := models.GetTagAccroingAritcle(1)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(ar)
	beego.Run()
}