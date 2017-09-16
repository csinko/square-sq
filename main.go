package main

import (//"fmt"
       //"html"
       "log"
       "net/http"
       )

func main() {
	
	fs := http.FileServer(http.Dir("dashboard"))
  	http.Handle("/dashboard/", http.StripPrefix("/dashboard", fs))

  	log.Println("Listening...")
  	http.ListenAndServe(":8080", nil)
    
}

/**
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}




*/