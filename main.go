package main

import (
	//"encoding/json"
	//"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
	//"strings"
	"html"
	"net/url"
	"net"
    )

//initialize an array of strings
var results []string

//function to collect post requests
func parsePost (w http.ResponseWriter, r *http.Request){



	fmt.Println(r.Method)

	if r.Method == "POST" {

		//fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path) 
		//url := html.EscapeString(r.URL.Path)
		//fmt.Println(url)
		//s := strings.Split(url,"id")

		//before := s[0]
		//after := s[1]
		//fmt.Println(before,after)
	
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{

			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		results = append(results, string(body))

		fmt.Fprint(w, "POST done")

	}else{

			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
}

func main() {
	

	//accept front end files
	fs := http.FileServer(http.Dir("dashboard"))
  	http.Handle("/dashboard/", http.StripPrefix("/dashboard", fs))


  	//create mux to identify post type
  	mux := http.NewServeMux()
  	mux.HandleFunc("/push", parsePost)

    /*s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
	*/



 
  	log.Println("Listening...")
   	http.ListenAndServe(":8080", mux)
}

