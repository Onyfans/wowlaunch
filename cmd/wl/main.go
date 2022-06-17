package main

import (
	"fmt"
	"github.com/KarazhanChessClub/wowlaunch/pkg/wowlaunch"
	"log"
	"os"
)

func printHelp() {
	fmt.Println(`
Usage: wowlaunch AccountName
AccountName must match a name variable in config.yaml`)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Print("Invalid number of arguments")
		printHelp()
		os.Exit(1)
	}

	c, err := wowlaunch.LoadConfigFromFile()
	checkErr(err)

	a, err := wowlaunch.ValidateAccount(c, os.Args[1])
	checkErr(err)

	err = wowlaunch.SetupRealmlist(c, a)
	checkErr(err)

	err = wowlaunch.StartWow(c)
	checkErr(err)

	wowlaunch.LoginScreenTyper(c, a)
}
