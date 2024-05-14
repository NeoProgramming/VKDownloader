package core

import (
	"github.com/BurntSushi/toml"
	"os"
	"vkdownloader/cfmt"
)

type Configuration struct {
	AppID       string
	AccessToken string
	RecentIP    string
	SavePath    string
}

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &App.config); err != nil {
		cfmt.PrintlnErr("Error reading configuration file:", err)
		// set default values
		// ...
	} else {
		cfmt.PrintlnOk("Configuration values loaded from config.toml")
	}
}

func SaveConfig() {
	file, err := os.Create("config.toml")
	if err != nil {
		cfmt.PrintlnErr("Error saving configuration file:", err)
		return
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(App.config); err != nil {
		cfmt.PrintlnErr(err)
		return
	}

	cfmt.PrintlnOk("Configuration values saved to config.toml")
}
