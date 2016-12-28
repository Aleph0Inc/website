package main

import(
    "time"
)

type Contact struct {
  ID int64 `json:"id"`
  Name string `json:"name"`
  Text string `json:"text" sql:"type:text;"`
  Email string `json:"email"`
  Created_at time.Time `json:"created_at"`
  Updated_at time.Time `json:"updated_at"`
}

type Contacts []Contact
func ContactCreate (name string, text string, email string){
  contact := Contact{Name: name, Text: text, Email: email,Created_at: time.Now(), Updated_at: time.Now()}

  db.NewRecord(&contact)
  db.Create(&contact)
  return
}

