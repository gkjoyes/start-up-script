package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

//InputValues provides struct for unmarshalling request info
type InputValues struct {
	Jobs []struct {
		Name         string   `json:"name"`
		Commands     []string `json:"commands"`
		Location     string   `json:"location,omitempty"`
		GitPull      bool     `json:"git_pull"`
		SourceBranch string   `json:"source_branch,omitempty"`
		DestBranch   string   `json:"dest_branch,omitempty"`
		Loaction     string   `json:"loaction,omitempty"`
	} `json:"jobs"`
}

func main() {
	//---------color output pointers
	info := color.New(color.FgGreen).Add(color.Bold)
	msg := color.New(color.FgHiCyan).Add(color.Bold)
	errMsg := color.New(color.FgRed).Add(color.Bold)

	//----------read json input file
	reqData, err := os.Open("conf/conf.json")
	if err != nil {
		errMsg.Println(err.Error())
		os.Exit(1)
	}
	//---------decoder
	reqParser := json.NewDecoder(reqData)
	reqJSON := InputValues{}
	err = reqParser.Decode(&reqJSON)
	if err != nil {
		errMsg.Println(err.Error())
		os.Exit(1)
	}
	//--------processing each jobs
	for _, job := range reqJSON.Jobs {
		var jobRun string
		for strings.TrimSpace(jobRun) != "n" && strings.TrimSpace(jobRun) != "y" {
			//----------message
			msg.Printf("Do you want to start ")
			info.Printf(job.Name)
			msg.Printf("[y/n]:")
			//---------user input
			_, err := fmt.Scan(&jobRun)
			if err != nil {
				errMsg.Println(err.Error())
			}
			//--------user input validation
			if strings.TrimSpace(jobRun) != "y" && strings.TrimSpace(jobRun) != "n" {
				errMsg.Println("invalid input! please enter valid input!")
			}
		}
		if jobRun == "y" {
			//----------create execution command
			cmd := "guake"
			execmds := "cd " + job.Loaction

			//----------git pull
			if job.GitPull {
				gitCommand := "git pull "
				if job.DestBranch != "" {
					gitCommand = gitCommand + " origin " + job.DestBranch
				}
				execmds = execmds + " && " + gitCommand
			}
			//---------other commands
			for _, execmd := range job.Commands {
				execmds = execmds + " && " + execmd
			}

			//----------pass arguments
			args := []string{"-n", " ", "-r", job.Name, "-e", execmds}
			if err = exec.Command(cmd, args...).Run(); err != nil {
				errMsg.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			info.Println("done")
		}
	}
}
