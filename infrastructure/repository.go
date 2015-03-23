package infrastructure

import (
  //"time"
    "fmt"
  	"github.com/videumcodeup/project-registration-system-go/infrastructure/entities"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/logs"
    	"github.com/astaxie/beego"
)

func ConfigureOrm(){
  	orm.RegisterDriver("mysql", orm.DR_MySQL)
  	orm.RegisterDataBase("default", "mysql", "@/codeup_reg_dev?charset=utf8")
  	orm.RegisterModel(new(entities.User), new(entities.Registration), new(entities.CodeUp))

    // Database alias.
    name := "default"
    // Drop table and re-create.
    force := false
    // Print log.
    verbose := true
    // Error.
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
      //log.Trace("%+v\n", err)
    }
}

type Repository interface {
    CreateUser(user entities.User) entities.User
    GetUserById(id int) entities.User
}


func NewApplicationRepository() *applicationRepository {
  ar := new(applicationRepository)
  ar.db = orm.NewOrm()

  return ar
}

type applicationRepository struct {
  db    orm.Ormer
}

func (this *applicationRepository) CreateUser(user entities.User) entities.User {
  beego.Trace("CreateUser...")
  id, err := this.db.Insert(&user)

  log := logs.NewLogger(10000)
  log.SetLogger("console", "")

  log.Trace("%+v\n", id)

  log.Trace("%+v\n", err)
  if err == nil {
    log.Trace("%+v\n", user)
  }
  return user
}
func (this *applicationRepository) CreateCodeUp(codeUp entities.CodeUp) entities.CodeUp{
  id, err := this.db.Insert(&codeUp)

  log := logs.NewLogger(10000)
  log.SetLogger("console", "")

  log.Trace("%+v\n", id)

  log.Trace("%+v\n", err)
  if err == nil {
    log.Trace("%+v\n", codeUp)
  }

  return codeUp
}


func (this *applicationRepository) GetUserById(id int) entities.User {
  user := entities.User{Id: id}
  err := this.db.Read(&user)

  if err == orm.ErrNoRows {
      fmt.Println("No result found.")
  } else if err == orm.ErrMissPK {
      fmt.Println("No primary key found.")
  } else {
      fmt.Println(user.Id, user.Name)
  }
  return user
}

func (this *applicationRepository) GetAllCodeUps() *[]entities.CodeUp{
  var codeUps []entities.CodeUp
  _, err := this.db.QueryTable(new(entities.CodeUp)).All(&codeUps)
  log := logs.NewLogger(10000)
  log.SetLogger("console", "")

  log.Trace("Err: %v",err)
    log.Trace("Result: %v",&codeUps)

  if err == orm.ErrNoRows {
      fmt.Println("No result found.")
  } else if err == orm.ErrMissPK {
      fmt.Println("No primary key found.")
  } else {
    //fmt.println(num)
  }
  return &codeUps
}
