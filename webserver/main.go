package main

import (
  "net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
  bar := "bar"

  w.Write([]byte("foo:" + bar))
}

func main() {
  http.HandleFunc("/status", status)

  if err := http.ListenAndServe(":2633", nil); err != nil {
    panic(err)
  }
}
