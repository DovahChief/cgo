package main

import "net/http"

func main()  {
    //http://localhost:8080/public/ej1/
    http.ListenAndServe(":8081", http.FileServer(http.Dir("")))
}
