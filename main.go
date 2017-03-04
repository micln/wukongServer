package main

import (
	_ "wukongServer/routers"

	"github.com/astaxie/beego"
)

func init() {
}

func main() {

	//wk := wukong.NewWukong()

	//
	//fmt.Println(wk.ToJson())
	//
	//fmt.Println(wk.SearchText(`百度中国`))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
