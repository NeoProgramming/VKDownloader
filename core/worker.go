package core

import (
	"time"
	"vkdownloader/cfmt"
)

const VkCount = 50

func (app *Application) worker() {
	defer app.wg.Done()
	// loop until the execution flag is cleared
	cfmt.PrintlnFunc("Worker running...")
	for app.running {
		// download
		time.Sleep(1 * time.Second)
	}
	app.running = false
	cfmt.PrintlnImp("Worker stopped naturally")
}
