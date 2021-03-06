package main

import (
	_ "suyan-service-catalog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}
	// For Beego CORS
	var corsFilter = cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"HEAD", "GET", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	beego.InsertFilter("/*", beego.BeforeRouter, corsFilter)
	beego.Run()
}
