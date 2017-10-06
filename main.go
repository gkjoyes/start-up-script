package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	info := color.New(color.FgGreen).Add(color.Bold)
	errMsg := color.New(color.FgRed).Add(color.Bold)
	info.Println("#-------------------startup-script--------------------#")
	var (
		// dashboard string
		genie string
		// airflow    bool
		// etlCodeM4m bool
		// dbScript   bool
		// Jobs       bool
	)
	//-----------genie
	fmt.Printf("Do you want to start dashboard[y/n] :")
	_, err := fmt.Scan(&genie)
	if err != nil {
		errMsg.Println(err.Error())
	}
	if genie != "y" || genie != "n" {
		errMsg.Println("invalid input!")
		os.Exit(0)
	}
}
