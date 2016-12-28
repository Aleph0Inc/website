package main

import "time"

type Todo struct { 
  Id int
  Name string
  Completed bool
  Due time.Time
}

type Todos []Todo
