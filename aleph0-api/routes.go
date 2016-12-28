package main

import (
  "net/http"

)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route


var routees = Routes{
Route{
    "CreateContact",
    "POST",
    "/contacts",
    CreateContact,
},
Route{
    "CreateSignup",
    "POST",
    "/signup",
    CreateSignup,
},
}
var routes = Routes{
//  Route{
//    "Index",
//    "GET",
//    "/",
//    Index,

//  },
  Route{
    "TodoIndex",
    "GET",
    "/todos",
    TodoIndex,

  },
  Route{
    "TodoShow",
    "GET",
    "/todos/{todoId}",
    TodoShow,

  },
Route{
    "TodoCreate",
    "POST",
    "/todos",
    TodoCreate,
},
Route{
    "UserCreate",
    "POST",
    "/signup",
    CreateUser,
},
Route{
    "LoginUser",
    "POST",
    "/signin",
    LoginUser,
},
Route{
    "CreateItem",
    "POST",
    "/items",
    ItemCreate,
},
Route{
    "ItemIndex",
    "GET",
    "/items",
    ItemIndex,
},
Route{
    "LogoutUser",
    "DELETE",
    "/signout",
    LogoutUser,
},
Route{
    "UpdateItem",
    "PUT",
    "/items/{id}",
    UpdateItem,
},
Route{
    "RemoveItem",
    "DELETE",
    "/items/{id}",
    RemoveItem,
},
}



