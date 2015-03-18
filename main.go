package main

import (
	_ "github.com/videumcodeup/project-registration-system-go/routers"
	"github.com/astaxie/beego"
	"github.com/videumcodeup/project-registration-system-go/infrastructure/entities"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "@/codeup_reg_dev?charset=utf8")
	orm.RegisterModel(new(entities.User))
}
