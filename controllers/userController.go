package controllers

import (
	"github.com/astaxie/beego"
	"github.com/videumcodeup/project-registration-system-go/infrastructure"
	"github.com/videumcodeup/project-registration-system-go/infrastructure/entities"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
  this.Data["Username"] = this.Ctx.Input.Param(":user")

  var repo = infrastructure.NewApplicationRepository()
  var user entities.User
  user.Name = "Senästien"
  user.Email = "sl@mail.com"
  repo.CreateUser(user)

  this.TplNames = "user.tpl"



}
