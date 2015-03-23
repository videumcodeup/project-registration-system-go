package main

import (
	_ "github.com/videumcodeup/project-registration-system-go/routers"
	"github.com/astaxie/beego"
	"github.com/videumcodeup/project-registration-system-go/infrastructure"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//"time"
	//"github.com/videumcodeup/project-registration-system-go/infrastructure/entities"


)

func main() {
	beego.Run()
}

func init() {
	infrastructure.ConfigureOrm()
	orm.Debug = true
	//fakedata.seedDb()
/*
	var repo = infrastructure.NewApplicationRepository()
	var c1 entities.CodeUp
	c1.Title = "CodeUp-träff 1"
	c1.Description = "Vi träffas, kodar, äter pizza och dricker öl. Mys ^^"
	c1.StartTime = time.Now().Add(time.Duration(48)*time.Hour)
	c1.EndTime = time.Now().Add(time.Duration(51)*time.Hour)

	var c2 entities.CodeUp
	c2.Title = "CodeUp-träff 2"
	c2.Description = "Vi träffas, kodar, äter pizza och dricker öl. Mys ^^"
	c2.StartTime = time.Now().Add(time.Duration(56)*time.Hour)
	c2.EndTime = time.Now().Add(time.Duration(59)*time.Hour)

	var c3 entities.CodeUp
	c3.Title = "Arduino-workshop"
	c3.Description = "Vi träffas, labbar med arduino, äter pizza och dricker öl. Mys ^^"
	c3.StartTime = time.Now().Add(time.Duration(64)*time.Hour)
	c3.EndTime = time.Now().Add(time.Duration(70)*time.Hour)

	c1 = repo.CreateCodeUp(c1)
	c2 = repo.CreateCodeUp(c2)
	c3 = repo.CreateCodeUp(c3)
*/
	beego.SetStaticPath("/img","static/img")
	beego.SetStaticPath("/js","static/js")
	beego.SetStaticPath("/css","static/css")
}
