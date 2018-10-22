## start_up_script

### Descriptions 

Automatically open multiple jobs in multiple windows of `guake` by a single click.

### Requirement 

* [Golang](https://golang.org/)

* [Guake](http://guake.org/)

### Usage 

* Provide your job details by using `conf/conf.json`.

* Example of config file show below.

* You can keep config file anywhere and just pass location of the file by using command line interface.

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

### Working 

* Run command `go run main.go` from project folder.
