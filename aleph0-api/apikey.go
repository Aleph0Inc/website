package main

import(
    "time"
)

type Apikey struct { 
  ID int64 `json:"id"`
  Owner string `json:"owner"`
  Secret string `json:"secret"`
  Created_at time.Time `json:"created_at"`
  Updated_at time.Time `json:"updated_at"`
}

type Apikeys []Apikey



func ApikeyCreate(owner string){

  db.DB()
  secret := randSeq(17)
  apikey := Apikey{Owner: owner, Secret: secret, Created_at: time.Now(), Updated_at: time.Now()}
  db.NewRecord(&apikey)
  db.Create(&apikey)
}

