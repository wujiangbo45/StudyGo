package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test/:id:int/", &controllers.TestController{})
}
