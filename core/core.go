package core

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"net/http"
	"os"
	"sync"
	"vkdownloader/cfmt"
)

type Application struct {
	config Configuration
	
	srv *http.Server
	vk  *api.VK
	
	totalItems      int
	currentItem     int
	running         bool

	wg sync.WaitGroup
}

var App Application

func InitCore() {

	cfmt.PrintlnImp("Starting VKExplorer app...")

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cfmt.PrintlnArg("Executable path:", exePath)
	info, err := os.Stat(exePath)
	if err == nil {
		cfmt.PrintlnArg("Build time:", info.ModTime())
	}
	curPath, err := os.Getwd()
	if err != nil {
		cfmt.PrintlnErr("Getwd error:", err)
	}
	cfmt.PrintlnArg("Current path:", curPath)

	LoadConfig()
}

func StartCore() {
	InitWeb()
	cfmt.PrintlnOk("Server is listening http://127.0.0.1:8080 ...")
}

func QuitCore() {
	cfmt.PrintlnOk("VKDownloader app finished")
}
