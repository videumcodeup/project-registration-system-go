package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
  this.Data["Username"] = this.Ctx.Input.Param(":user")
	beego.Info(this.Data["Username"])
	/*
  var repo = infrastructure.NewApplicationRepository()
  var user entities.User
  user.Name = "Senästien"
  user.Email = "sl@mail.com"
  repo.CreateUser(user)
	*/
	this.Layout = "layout.html"
  this.TplNames = "user.tpl"
}
