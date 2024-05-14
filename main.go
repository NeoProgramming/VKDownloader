package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"vkdownloader/cfmt"
	"vkdownloader/core"
)

func main() {
	core.InitCore()
	core.StartCore()

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptCh
		cfmt.PrintlnImp("\nReceived interrupt signal. Exiting...")
		core.QuitCore()
		os.Exit(0)
	}()

	for {
		time.Sleep(time.Second)
	}
}
