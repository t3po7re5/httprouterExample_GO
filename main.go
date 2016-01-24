package main

import (
  "strconv"
  "fmt"
  "log"
  "net/http"
)

func main() {
  var PORT_NUMBER = 8080
  router := NewRouter()
  fmt.Println("Server is listening on port:", PORT_NUMBER)
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT_NUMBER), router))
}
