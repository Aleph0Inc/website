package main

import(
    "time"
)

type Item struct {
  ID int64 `json:"id"`
  UserID int64 `json:"user_id" sql:"index"`
  Name string `json:"name"`
  Desc string `json:"desc"`
  Lat float64 `json:"lat"`
  Lon float64  `json:"lon"`
  Created_at time.Time `json:"created_at"`
  Updated_at time.Time `json:"updated_at"`
}

type Items []Item
func CreateItem (u *User, name string, desc string, lat float64, lon float64) (i *Item){
  item := Item{UserID: u.ID, Name: name, Desc: desc, Lat: lat, Lon: lon, Created_at: time.Now(), Updated_at: time.Now()}

  i = &item
  db.NewRecord(&item)
  db.Create(&item)
  return
}


