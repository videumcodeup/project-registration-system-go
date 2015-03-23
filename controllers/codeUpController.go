package controllers

import (
	"github.com/astaxie/beego"
	"github.com/videumcodeup/project-registration-system-go/infrastructure"
	"github.com/videumcodeup/project-registration-system-go/common/helpers"
)

type CodeUpController struct {
	beego.Controller
}

func (this *CodeUpController) Get() {
	this.Layout = "layout.html"
  this.TplNames = "codeups.tpl"
}

func (this *CodeUpController) GetAll(){
	var repo = infrastructure.NewApplicationRepository()
	var codeUps = repo.GetAllCodeUps()

	var leng = len(codeUps)

	/**
	/ The JSON-object returned will be null if the key
	/ is not "json" in the Data-dict.
	/ see http://beego.me/docs/mvc/controller/jsonxml.md
	**/
	//this.Data["json"] = modelhelper.MapToCodeUpViewModelList(&codeUps)
  this.ServeJson()
}
