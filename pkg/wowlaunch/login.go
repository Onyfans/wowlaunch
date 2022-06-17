package wowlaunch

import (
	"fmt"
	rg "github.com/go-vgo/robotgo"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func StartWow(c *Config) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command(path.Join(c.Wowpath, "Wow.exe"))
	} else {
		cmd = exec.Command("wine", path.Join(c.Wowpath, "Wow.exe"))
		cmd.Env = append(os.Environ(), fmt.Sprintf("WINEPREFIX=%s", c.Wineprefix))
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("could not start wow. %v \n", err)
	}
	return nil
}

func LoginScreenTyper(c *Config, a *Account) {
	rg.Sleep(c.Delay)
	rg.TypeStr(a.Username)
	rg.KeyTap("tab")
	rg.TypeStr(a.Password)
	rg.KeyTap("enter")
}
