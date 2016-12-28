package main

import(
    "time"
)

type Signup struct {
  ID int64 `json:"id"`
  Email string `json:"email"`
  Created_at time.Time `json:"created_at"`
  Updated_at time.Time `json:"updated_at"`
}

type Signups []Signup
func SignupCreate (email string){
  signup := Signup{Email: email,Created_at: time.Now(), Updated_at: time.Now()}

  db.NewRecord(&signup)
  db.Create(&signup)
  return
}

