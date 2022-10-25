package routers

import (
	"burlyeducation/controllers"
	v1 "burlyeducation/controllers/v1"
	"burlyeducation/lib"

	beego "github.com/beego/beego/v2/server/web"
	//"github.com/beego/admin"
)

// func init() {
// 	//admin.Run()

// }

func init() {
	beego.Router("/", &controllers.MainController{})

	authUrls := []string{"/backend-burlyed/v1/users"}

	for _,url := range authUrls {
		beego.InsertFilter(url, beego.BeforeRouter, lib.ApplyAuth)
	}
	
	beego.ErrorController(&controllers.ErrorController{})
	
	ns := beego.NewNamespace("/burlyed/v1/",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&v1.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/backend-burlyed/v1/user", &v1.UserController{})
}
