package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func PewRouter() *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routees {
    var handler http.Handler
    handler = route.HandlerFunc
    handler = Logger(handler, route.Name)

    router.
      Methods(route.Method).
      Path("/api/voo"+route.Pattern).
      Name(route.Name).
      Handler(handler)
  }
  return router
}

