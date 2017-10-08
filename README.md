**project_name** : start_up_script
***
#**description** : This is simple project for open multiple `guake` terminal window with different locations and execute different command on each window.
***
#**usage** : add `jobs` entry in json file
      
                  {
                        "name": "project_name",
                        "commands": [
                            "command need to execute"
                        ],
                        "loaction": "project_loaction",
                        "git_pull": true,
                        "dest_branch": "which branch you need to sync with"
                   }

