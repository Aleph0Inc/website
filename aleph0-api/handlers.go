package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv" 
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
//    todos := Todos{
//        Todo{Name: "Write presentation"},
//        Todo{Name: "Host meetup"},
//    }



    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      panic(err)
    }
    if err := r.Body.Close(); err != nil {
      panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(422) //unprocessable entity
      if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
      }
    }

  t := RepoCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(t); err != nil {
    panic(err)
  }
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
  user := UserCreate(r.FormValue("username"), r.FormValue("password"))
  session, _ := store.Get(r, "Aleph0-session")
  session.Values["userid"] = user.ID
  session.Values["loggedin"] = "true"
  session.Save(r, w)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if user != nil {
    if err := json.NewEncoder(w).Encode(user); err != nil {
      panic(err)
    }
  }

}
func CreateContact(w http.ResponseWriter, r *http.Request) {
  ContactCreate(r.FormValue("name"), r.FormValue("text"), r.FormValue("email"))
  http.Redirect(w, r, "/index.htm", 301)
}
func CreateSignup(w http.ResponseWriter, r *http.Request) {
  SignupCreate(r.FormValue("email"))
  http.Redirect(w, r, "/index.htm", 301)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
  u, err := Login(r.FormValue("username"), r.FormValue("password"))
  if err != nil{
    panic(err)
  }
  if u != nil {
    session, _ := store.Get(r, "Aleph0-session")
    session.Values["userid"] = u.ID
    session.Values["loggedin"] = "true"
    session.Save(r, w) 
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(u); err != nil {
      panic(err)
    }
  }

}
func LogoutUser(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "Aleph0-session")
    session.Values["userid"] = 0
    session.Values["loggedin"] = "false"
    session.Save(r, w) 
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}
func ItemIndex(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "Aleph0-session")
    id := session.Values["userid"]
    loggedin := session.Values["loggedin"]
    if id != nil && id != 0 && loggedin == "true" {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      items := Items{}
      db.Where("user_id = ?", id).Find(&items)
      if err := json.NewEncoder(w).Encode(items); err != nil {
          panic(err)
      }
    }

}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "Aleph0-session")
    id := session.Values["userid"]
    loggedin := session.Values["loggedin"]
    vars := mux.Vars(r)
    itemid := vars["id"]
    item := Item{}
    db.Find(&item, itemid)
    lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
    if err != nil {
      panic(err)
    }
    lon, err := strconv.ParseFloat(r.FormValue("lon"), 64)
    if err != nil {
      panic(err)
    }
//    userid, err := strconv.ParseInt(vars["id"], 0, 0)
//    if err != nil {
//      panic(err)
//    }
    if id != nil && id != 0 && loggedin == "true" && item.UserID == id {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      item.Desc = r.FormValue("desc")
      item.Name = r.FormValue("name")
      item.Lat = lat
      item.Lon = lon
      db.Save(&item)
      if err := json.NewEncoder(w).Encode(item); err != nil {
       panic(err)
       }
    }
}
func RemoveItem(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "Aleph0-session")
    id := session.Values["userid"]
    loggedin := session.Values["loggedin"]
    vars := mux.Vars(r)
    itemid := vars["id"]
    item := Item{}
    db.Find(&item, itemid)
//    userid, err := strconv.ParseInt(vars["id"], 0, 0)
//    if err != nil {
//      panic(err)
//    }
    if id != nil && id != 0 && loggedin == "true" && item.UserID == id {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      db.Delete(&item)
    }
}
func ItemCreate(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "Aleph0-session")
  id := session.Values["userid"]
  loggedin := session.Values["loggedin"]
  lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
  if err != nil {
    panic(err)
  }
  lon, err := strconv.ParseFloat(r.FormValue("lon"), 64)
  if err != nil {
    panic(err)
  }
  if id != nil && id != 0 {
    if loggedin == "true" {
    user := User{}
    db.Find(&user, id)
    item := CreateItem(&user, r.FormValue("name"), r.FormValue("desc"), lat, lon)
     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
     w.WriteHeader(http.StatusOK)
     if item != nil {
       if err := json.NewEncoder(w).Encode(item); err != nil {
       panic(err)
       }
     }
   }
  }
}




