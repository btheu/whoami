package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "strings"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    name := os.Getenv("NAME")
    desc := os.Getenv("DESC")
    ctxt := os.Getenv("CONTEXT_PATH")

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/" + ctxt, func(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(os.Stdout, "I'm %s \n", hostname)
       fmt.Fprintf(os.Stdout, "Name: %s \n", name)
       fmt.Fprintf(os.Stdout, "Desc: %s \n", desc)
       
       fmt.Fprintf(w, "I'm %s \n", hostname)
       fmt.Fprintf(w, "Name: %s \n", name)
       fmt.Fprintf(w, "Desc: %s \n", desc)
       
       fmt.Fprintf(os.Stdout, "==Env \n")
       fmt.Fprintf(w, "==Env \n")
       
       for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            fmt.Fprintf(os.Stdout, "%s = %s \n", pair[0], pair[1])
            fmt.Fprintf(w,         "%s = %s \n", pair[0], pair[1])
        }

       fmt.Fprintf(os.Stdout, "==Headers \n")
       fmt.Fprintf(w,         "==Headers \n")

       for name, values := range r.Header {
           for _, value := range values {
             fmt.Fprintf(os.Stdout, "%s = %s \n", name, value)
             fmt.Fprintf(w,         "%s = %s \n", name, value)
           }
        }

    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

