**project_name** : `start_up_script`
***
**description** : This is simple project for open multiple `guake` terminal window with different locations and execute different command on each window.
***
**requirement** : `golang`
***
**usage** : add `jobs` entry in json file


      ```{
          "jobs": [
              {
                  "name": "job_1",
                  "commands": [
                      "echo job_1 starting.....",
                      "echo job_1 done...."
                  ],
                  "loaction": "/home/georgekutty/work/javascript",
                  "git_pull": true,
                  "dest_branch": "sub_branch"
              },
              {
                  "name": "job_2",
                  "commands": [
                      "echo job_2 starting.....",
                      "echo job_2 done...."
                  ],
                  "loaction": "/home/georgekutty/work/javascript",
                  "git_pull": false,
                  "dest_branch": ""
              }
          ]
      }```



***
**run-tool** : run command `go run main.go` from project folder
