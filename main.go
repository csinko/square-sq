package main

import (
	"encoding/json"
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
		
			// Unmarshal
			var msg Webhook
			err = json.Unmarshal(body, &msg)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			output, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("content-type", "application/json")
			w.Write(output)

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

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
