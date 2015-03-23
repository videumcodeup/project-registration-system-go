package entities

import (

)

type User struct {
    Id    int
    Name  string
    Email string `orm:"unique"`
    AccessLevel int16
    Registrations []*Registration `orm:"reverse(many)"`
}
