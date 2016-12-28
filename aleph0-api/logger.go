package main

import (
  "net/http"
  "log"
  "time"
  "os"
  "io"
)

func Logger(inner http.Handler, name string) http.Handler {

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    inner.ServeHTTP(w, r)

    log.Printf("%s\t%s\t%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
      r.RemoteAddr,
      r.Header.Get("X-FORWARDED-FOR"),
    )
    f, err := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
      panic(err)
    }
    defer f.Close()
    log.SetOutput(f)
    mylogger := log.New(io.MultiWriter(f,os.Stdout), "", 0)
    log.Printf("%s\t%s\t%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
      r.RemoteAddr,
      r.Header.Get("X-FORWARDED-FOR"),
    )
    mylogger.Printf("%s\t%s\t%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
      r.RemoteAddr,
      r.Header.Get("X-FORWARDED-FOR"),
    )
  })

}
