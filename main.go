package main

import (
	"errors"
	"fmt"
	rg "github.com/go-vgo/robotgo"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type Config struct {
	Delay      int
	Wowpath    string
	Wineprefix string
	Accounts   []Account
}

type Account struct {
	Name     string
	Username string
	Password string
}

func init() {
	viper.SetDefault("delay", 5)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	p, _ := os.Executable()
	viper.AddConfigPath(".")
	viper.AddConfigPath(path.Dir(p))
	viper.AddConfigPath(`.config/wowlaunch/`)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v \n", err)
	}
	if !viper.IsSet("wowpath") {
		log.Fatal("fatal error: wowpath is not set in config.yaml")
	} else if _, err := os.Stat(viper.GetString("wowpath")); os.IsNotExist(err) {
		log.Fatal("fatal error: could not find Wow.exe, check wowpath in config.yaml")
	}
}

func validateAccount(c *Config, s string) (int, error) {
	for i, ac := range c.Accounts {
		if s == ac.Name {
			return i, nil
		}
	}
	return 0, errors.New("provided name was not found in config.yaml")
}

func printHelp() {
	fmt.Println(`
Usage: wowlaunch AccountName
AccountName must match a name variable in config.yaml`)
}

func main() {
	if len(os.Args) != 2 {
		log.Print("Invalid number of arguments")
		printHelp()
		os.Exit(1)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		log.Print("fatal error: could not unmarshal config.yaml")
	}

	idx, err := validateAccount(&c, os.Args[1])
	if err != nil {
		log.Fatalf("fatal error: %v \n", err)
	}

	cmd := exec.Command("wine", c.Wowpath)
	if runtime.GOOS != "windows" {
		cmd.Env = append(os.Environ(), fmt.Sprintf("WINEPREFIX=%s", c.Wineprefix))
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("fatal error: could not start wow. %v \n", err)
	}
	rg.Sleep(c.Delay)
	rg.TypeStr(c.Accounts[idx].Username)
	rg.KeyTap("tab")
	rg.TypeStr(c.Accounts[idx].Password)
	rg.KeyTap("enter")
}
