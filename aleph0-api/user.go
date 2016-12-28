package main

import (
  "code.google.com/p/go.crypto/bcrypt"
  "time"
)

type User struct {
  ID int64 `json:"id"`
  Username string `json:"username"`
  Password []byte `json:"none"`
  Posts int `json:"none"`
  Secret string `json:"none"`
  Created_at time.Time `json:"created_at"`
  Updated_at time.Time `json:"updated_at"`
}

type Users []User
func (u *User) SetPassword(password string) {
  hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    panic(err)
  }
  u.Password = hpass
}

//login validates and returns user object if it exists in database
func Login(username, password string) (u *User, err error) {
//  err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
//  if err != nil {
//    return
//  }
  user := User{}
  db.Where("username = ?", username).First(&user)
  u = &user
  err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
  if err != nil{
    u = nil
  }
  return

}


func UserCreate(username string, password string) (u *User){

  db.DB()

  user := User{Username: username, Created_at: time.Now(), Updated_at: time.Now()}
  user.SetPassword(password)
  luser := User{}
  d := User{}
  db.Where("username = ?", username).First(&luser)
  if (luser.Username != User{}.Username) {
    u = &d
    return
  }
  u = &user
  db.NewRecord(&user)
  db.Create(&user)
  return
}
