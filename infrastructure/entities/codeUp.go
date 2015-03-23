package entities

import (
  "time"
)

type CodeUp struct {
    Id    int
    Title string
    Description string
    StartTime time.Time
    EndTime time.Time
    Registrations []*Registration `orm:"reverse(many)"`
}
