package main

import (
	"io/ioutil"
	//"fmt"
	"log"
	"strings"
)

func UpdateConfig(appType string, user string, repo string) {
	//Open nginx config
	conf, err := ioutil.ReadFile("/etc/nginx/sites-available/default")
	if err != nil {
		panic(err)
	}

	//Split config into lines for easier handling
	confLines := strings.Split(string(conf), "\n")

	//Locate the last bracket (the closing server bracket)
	lastBracket := 0
	for i, line := range confLines {
		if strings.Contains(line, "}") {
			lastBracket = i
		}
	}
	//Panic if it isn't found
	if lastBracket == 0 {
		panic("Closing Bracket not found!")
	}

	var newConf []string

	newConf = append(newConf, "")
	newConf = append(newConf, "\t# START "+user+"/"+repo)

	locString := "\tlocation /repo/" + user + "/" + repo

	//TODO add compat to check if web app or not (assuming true for now)
	var isWeb bool
	if appType == "web" {
		isWeb = true
	}
	if !isWeb {
		locString += "/"
	}

	locString += " {"

	newConf = append(newConf, locString)

	//var isNode bool
	//if appType == "node" {
	//	isNode = true
	//}

	if isWeb {
		newConf = append(newConf, "\t\talias /var/app/deploy/"+user+"/"+repo+";")
	}

	newConf = append(newConf, "\t}")
	newConf = append(newConf, "\t# END "+user+"/"+repo)

	confLines = append(confLines[:lastBracket], append(newConf, confLines[lastBracket:]...)...)

	confOutput := strings.Join(confLines, "\n")
	err = ioutil.WriteFile("/etc/nginx/sites-available/default", []byte(confOutput), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}
