package main

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateApp(appType string, user string, repo string) {
	var err error

	rootAppFolder := "/var/app/deploy/" + user + "/" + repo
	os.MkdirAll(rootAppFolder, os.ModePerm)
	cmdName := "git"
	cmdArgs := []string{"clone", "https://github.com/" + user + "/" + repo, rootAppFolder}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Cloning Error: ", err)
		os.Exit(1)
	}

	UpdateConfig(appType, user, repo)

	cmdName = "systemctl"
	cmdArgs = []string{"restart", "nginx"}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Nginx Restart Error: ", err)
		os.Exit(1)
	}
}

//func UpdateApp(appType string, user string, repo string) {
//rootAppFolder := "/var/app/deploy/"+user+"/"+repo
//
//
//}
