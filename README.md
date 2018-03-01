## start_up_script



### Description 

 * This is simple project for open multiple `guake` terminal window with different locations and execute necessary command on each window


### requirement 

* [Golang](https://golang.org/)

* [Guake](http://guake.org/)


### usage 

* Update `conf/conf.json` depending on user demand

* Sample structure of `conf/conf.json` shown below


```
      {
          "jobs": [
              {
                  "name": "job_1",
                  "commands": [
                      "echo job_1 starting",
                      "echo job_1 done"
                  ],
                  "loaction": "/home/georgekutty/work/javascript",
                  "git_pull": true,
                  "dest_branch": "sub_branch"
              },
              {
                  "name": "job_2",
                  "commands": [
                      "echo job_2 starting",
                      "echo job_2 done"
                  ],
                  "loaction": "/home/georgekutty/work/javascript",
                  "git_pull": false,
                  "dest_branch": ""
              }
          ]
      }
```




### run-tool 

* run command `go run main.go` from project folder
