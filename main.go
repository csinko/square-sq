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
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getRepos(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		session := s.Copy()
		defer session.Close()

		c := session.DB("hackthenorth17").C("repos")

		var repos []Webhook
		err := c.Find(bson.M{}).All(&repos)

		if err != nil {
			http.Error(w, "Error failed to get repos", http.StatusInternalServerError)
			fmt.Println("Failed to get repos: ", err)
			return
		}

		respBody, err := json.MarshalIndent(repos, "", "  ")

		if err != nil {
			fmt.Println("Failed to convert")
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(respBody)

	}
}

func deleteRepos(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		session := s.Copy()
		defer session.Close()

		c := session.DB("hackthenorth17").C("repos")

		err := c.Remove(bson.M{})

		if err != nil {
			switch err {
			default:
				http.Error(w, "Database Error", http.StatusInternalServerError)

				fmt.Println("Failed to delete repo")
				return
			case mgo.ErrNotFound:
				http.Error(w, "Record Not Found", http.StatusNotFound)
				return

			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

//function to collect post requests
func parsePost(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		session := s.Copy()
		defer session.Close()

		c := session.DB("hackthenorth17").C("repos")

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

			var curRepo Webhook
			err = c.Find(bson.M{"repository.id": msg.Repository.ID}).One(&curRepo)

			//If the repo was not found (its new)
			if err != nil {
				if err = c.Insert(msg); err != nil {
					panic(err)
				}
				fmt.Println("Added to Database")
				fmt.Fprint(w, "Success")
				if msg.Repository.Language == "HTML" || msg.Repository.Language == "CSS" {
					CreateApp("web", msg.Repository.Owner.Name, msg.Repository.Name)
				}
				return
			}

			curRepo.Commits = append(curRepo.Commits, msg.Commits...)

			err = c.Update(bson.M{"repository.id": curRepo.Repository.ID}, &curRepo)
			if err != nil {
				fmt.Println("Failed to Update")
				fmt.Fprint(w, "FAIL")
			}
			IsNodeApplication(msg)
			fmt.Println("Testing Complete")

			if msg.Repository.Language == "HTML" || msg.Repository.Language == "CSS" {
				UpdateApp("web", msg.Repository.Owner.Name, msg.Repository.Name)
			}

			fmt.Println("Updated")
			fmt.Fprint(w, "Success")

			//Send success as a response

		} else {
			//Return an error if not POST
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

func main() {

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"ds139844.mlab.com:39844"},
		Username: "htn17",
		Password: "htn2017",
		Database: "hackthenorth17",
	})
	if err != nil {
		panic(err)
	}

	defer session.Close()
	fmt.Printf("Connected to %v!\n", session.LiveServers())

	mux := mux.NewRouter()
	mux.HandleFunc("/push", parsePost(session))
	mux.HandleFunc("/api/repos", getRepos(session))
	mux.HandleFunc("/api/repos/delete", deleteRepos(session))
	fs := http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./dashboard/")))
	mux.PathPrefix("/dashboard/").Handler(fs)
	http.Handle("/", mux)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
