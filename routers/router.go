package routers

import (
	"github.com/videumcodeup/project-registration-system-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/Hej", &controllers.MainController{})
		beego.Router("/User/:user", &controllers.UserController{})
		beego.Router("/CodeUps", &controllers.CodeUpController{})
		beego.Router("/CodeUps/GetAll", &controllers.CodeUpController{}, "get:GetAll")
}
