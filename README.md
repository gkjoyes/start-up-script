## start_up_script


&nbsp;
### Description 
***

Automatically open multiple jobs in multiple windows of `guake` in a single click.

&nbsp;
### Requirement 
***

* [Golang](https://golang.org/)

* [Guake](http://guake.org/)

&nbsp;
### Usage 

* Update `conf/conf.json` depending on user demand

* Sample structure of `conf/conf.json` shown below


```
{
    "jobs": [
        {
            "name": "job-1",
            "commands": [
                "echo starting",
                "echo done"
            ],
            "path": "/home/go"
        },
        {
            "name": "job-2",
            "commands": [
                "echo starting",
                "echo done"
            ],
            "path": "/home/python"
        }
    ]
}
```

&nbsp;
### Working 
* Run command `go run main.go` from project folder.