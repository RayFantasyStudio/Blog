package main

import (
	_ "github.com/RayFantasyStudio/blog/routers"
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/utils"
	"github.com/RayFantasyStudio/blog/models"
)
func init(){
	var err error
	_,models.TotalArticleCount,err =models.GetArticles(models.Filter_Create,"","",false)
	if err != nil {
		beego.Error(err)
	}
	beego.Info("Total article count = ",models.TotalArticleCount)
}

func main() {

	beego.AddFuncMap("SinceTime",utils.SinceTime)
	beego.Run()
}