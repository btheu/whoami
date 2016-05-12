package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    name := os.Getenv("NAME")
    desc := os.Getenv("DESC")

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s \n", hostname)
        fmt.Fprintf(os.Stdout, "Name: %s \n", name)
        fmt.Fprintf(os.Stdout, "Desc: %s \n", desc)
        
 	      fmt.Fprintf(w, "I'm %s <br/>", hostname)
 	      fmt.Fprintf(w, "Name: %s <br/>", name)
 	      fmt.Fprintf(w, "Desc: %s <br/>", desc)
 	      
 	      fmt.Fprintf(os.Stdout, "==Env \n")
        fmt.Fprintf(w, "==Env <br/>")
 	      
 	      for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            fmt.Fprintf(os.Stdout, "%s = %s \n", pair[0], pair[1])
            fmt.Fprintf(w, "%s = %s <br/>", pair[0], pair[1])
        }
 	      
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

