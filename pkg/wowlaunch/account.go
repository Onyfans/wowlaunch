package wowlaunch

import (
	"fmt"
)

type Account struct {
	Name     string
	Username string
	Password string
	Server   string
}

func ValidateAccount(c *Config, s string) (*Account, error) {
	for i, ac := range c.Accounts {
		if s == ac.Name {
			return &c.Accounts[i], nil
		}
	}
	return nil, fmt.Errorf("provided name was not found in config.yaml")
}
