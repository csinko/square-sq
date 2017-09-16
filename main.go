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
)

//function to collect post requests
func parsePost(w http.ResponseWriter, r *http.Request) {

	//Check if request is a POST request
	if r.Method == "POST" {

		//Read the body of the request.  If thers an error, return an error.
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		//Print out body for debug
		fmt.Println(string(body))

		//Parse body JSON into object

		//Send success as a response
		fmt.Fprint(w, "Success")

	} else {
		//Return an error if not POST
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
