package infrastructure

import (
    "fmt"
  	"github.com/videumcodeup/project-registration-system-go/infrastructure/entities"
    "github.com/astaxie/beego/orm"
)

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

func (ur applicationRepository) CreateUser(user entities.User) entities.User {
  id, err := ur.db.Insert(&user)
  if err == nil {
    fmt.Println(id)
  }
  return user
}

func (ur applicationRepository) GetUserById(id int) entities.User {
  user := entities.User{Id: id}
  err := ur.db.Read(&user)

  if err == orm.ErrNoRows {
      fmt.Println("No result found.")
  } else if err == orm.ErrMissPK {
      fmt.Println("No primary key found.")
  } else {
      fmt.Println(user.Id, user.Name)
  }
  return user
}

func (ur *applicationRepository) setOrm (o orm.Ormer){
  ur.db = o
}
