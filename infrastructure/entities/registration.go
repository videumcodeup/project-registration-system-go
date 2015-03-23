package entities

import (
  "time"
)

type Registration struct {
    Id    int
    Codeup  *CodeUp `orm:"rel(one);rel(fk)"`
    User  *User `orm:"rel(one);rel(fk)"`
    RegistrationTime  time.Time
    Message string
}

// multiple fields unique key
func (r *Registration) TableUnique() [][]string {
    return [][]string{
        []string{"CodeUp", "User"},
    }
}
