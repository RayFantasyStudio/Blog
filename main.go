package main

import (
	_ "github.com/RayFantasyStudio/blog/routers"
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/utils"
)

func main() {
	beego.AddFuncMap("SinceTime", utils.SinceTime)
	beego.Run()
}