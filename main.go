package main

import (
  "fmt"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Person struct {
  Name string
  Phone string
}
