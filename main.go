package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

const inputFile = "conf/conf.json"

//inputValues provides struct for unmarshalling request info
type inputValues struct {
	Jobs []struct {
		Name         string   `json:"name"`
		Commands     []string `json:"commands"`
		Location     string   `json:"location,omitempty"`
		GitPull      bool     `json:"git_pull"`
		SourceBranch string   `json:"source_branch,omitempty"`
		DestBranch   string   `json:"dest_branch,omitempty"`
	} `json:"jobs"`
}

func main() {

	// ----sys out colors
	errMsg := color.New(color.FgRed)

	// ----read json input file
	reqData, err := os.Open(inputFile)
	if err != nil {
		errMsg.Println("#Error: ", err.Error())
		os.Exit(0)
	}

	// ----decoder
	reqParser := json.NewDecoder(reqData)
	reqJSON := inputValues{}
	err = reqParser.Decode(&reqJSON)
	if err != nil {
		errMsg.Println("#Error: ", err.Error())
		os.Exit(0)
	}

	// ----validation
	err = validation(reqJSON)
	if err != nil {
		errMsg.Println("#Error: ", err.Error())
		os.Exit(0)
	}

	// ----running shell
	shell := "guake"

	// ----processing each jobs
	for _, job := range reqJSON.Jobs {

		// ----location
		execmds := "cd " + job.Location

		// ----git pull
		if job.GitPull {
			gitCommand := "git pull "
			if job.DestBranch != "" {
				gitCommand = gitCommand + " origin " + job.DestBranch
			}
			execmds = execmds + " && " + gitCommand
		}

		// ----other commands
		for _, execmd := range job.Commands {
			execmds = execmds + " && " + execmd
		}

		// ----pass arguments
		args := []string{"-n", " ", "-r", job.Name, "-e", execmds}

		// ----execution
		if err = exec.Command(shell, args...).Run(); err != nil {
			errMsg.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}
	}
}

// validation of user inputs
func validation(reqJSON inputValues) error {
	for _, job := range reqJSON.Jobs {
		// ----check project path is exists
		if _, err := os.Stat(job.Location); err != nil {
			return fmt.Errorf("job location %s is not exists", job.Location)
		}
	}
	return nil
}
