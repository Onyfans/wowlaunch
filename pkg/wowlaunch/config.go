package wowlaunch

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

type Config struct {
	Delay      int
	Wowpath    string
	Wineprefix string
	Region     string
	Servers    map[string]string
	Accounts   []Account
}

func LoadConfigFromFile() (*Config, error) {
	viper.SetDefault("delay", 5)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	p, _ := os.Executable()
	viper.AddConfigPath(".")
	viper.AddConfigPath(path.Dir(p))
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not find home dir. %v\n", err)
	}
	viper.AddConfigPath(path.Join(home, `.config/wowlaunch/`))

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if !viper.IsSet("wowpath") {
		return nil, fmt.Errorf("wowpath is not set in config.yaml")
	} else if _, err := os.Stat(viper.GetString("wowpath")); os.IsNotExist(err) {
		return nil, fmt.Errorf("could not find Wow.exe, check wowpath in config.yaml")
	}

	var c *Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("could not unmarshal config.yaml")
	}

	return c, nil
}

func SetupRealmlist(c *Config, a *Account) error {
	realmpath := path.Join(c.Wowpath, "Data", c.Region, "realmlist.wtf")
	realmfile, err := os.Create(realmpath)
	if err != nil {
		return fmt.Errorf("could not find realmlist.wtf. %v\n", err)
	}
	if _, ok := c.Servers[a.Server]; !ok {
		return fmt.Errorf("failed to match server in account list to server list")
	}
	_, err = realmfile.WriteString(fmt.Sprintf("set realmlist %s", c.Servers[a.Server]))
	return err
}
