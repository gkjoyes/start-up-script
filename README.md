**project_name** : `start_up_script`
***
**description** : This is simple project for open multiple `guake` terminal window with different locations and execute different command on each window.
***
**requirement** : `golang`
***
**usage** : add `jobs` entry in json file
      
                  {
                        "name": "project_name",
                        "commands": [
                            "command need to execute"
                        ],
                        "loaction": "project_loaction",
                        "git_pull": true,
                        "dest_branch": "which branch you need to sync with"
                   }
***
**run-tool** : run command `go run main.go` from project folder
