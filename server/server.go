package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  fmt.Fprintf(w, "server has responded :)\n")

}

func main() {

  router := httprouter.New()
  router.GET("/", homeHandler)

  fmt.Println("Starting server on port 8080")
  http.ListenAndServe(":8080", router)

}
