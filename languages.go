package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// IsNodeApplication checks if the app is a node app
func IsNodeApplication(msg Webhook) bool {
	//Get the last commit ID
	commitID := msg.Commits[len(msg.Commits)-1].ID

	//Get the root tree of the repository
	resp, err := http.Get("https://api.github.com/repos/" + msg.Repository.Owner.Name + "/" + msg.Repository.Name + "/git/trees/" + commitID)

	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var respTree GitTree

	if err := json.Unmarshal([]byte(respBody), &respTree); err != nil {
		log.Fatal(err)
	}

	for _, file := range respTree.Tree {
		if strings.Contains(file.Path, "package.json") {
			return true
		}
	}
	return false
}
