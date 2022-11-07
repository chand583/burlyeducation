package routers

import (
	"burlyeducation/controllers"
	v1 "burlyeducation/controllers/v1"

	"github.com/astaxie/beego"
	//"github.com/beego/admin"
)

func init() {
	//admin.Run()
	beego.Router("/", &controllers.MainController{})

	// authUrls := []string{"/backend-abcast/v1/article/create"}

	// for _, url := range authUrls {
	// 	beego.InsertFilter(url, beego.BeforeRouter, lib.ApplyAuth)
	// }

	// beego.ErrorController(&controllers.ErrorController{})

	// ns := beego.NewNamespace("/burlyed/v1/",
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&v1.UserController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
	//beego.Router("/backend-abcast/v1/article/create", &v1.ExamInfoController{})

	ns :=
		beego.NewNamespace("burlyed/v1",
			beego.NSRouter("/user", &v1.UserController{}),
			beego.NSRouter("/user/:id([0-9])+", &v1.UserController{}, "GET:GetOne"),
			// beego.NSRouter("/scheduler/weather",&controllers.ScheduleController{}),

		)
	beego.AddNamespace(ns)

}
