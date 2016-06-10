package main

import (
	_ "github.com/RayFantasyStudio/blog/routers"
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/utils"
	"github.com/RayFantasyStudio/blog/models"
	"github.com/astaxie/beego/orm"
)
func init(){
	var err error
	o := orm.NewOrm()
	models.TotalArticleCount,err = o.QueryTable("article").Count()
	if err != nil {
		beego.Error(err)
	}
	beego.Info("Total article count = ",models.TotalArticleCount)
}

func main() {

	beego.AddFuncMap("SinceTime",utils.SinceTime)
	beego.Run()
}