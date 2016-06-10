package models

import "github.com/astaxie/beego"

func init() {
	err := SetupTotalArticleCount()
	if err != nil {
		beego.Error(err)
	}
}