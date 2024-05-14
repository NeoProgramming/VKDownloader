package core

import (
	"log"
	"net/http"
	"vkdownloader/cfmt"
)

func InitWeb() {
	App.srv = &http.Server{
		Addr: ":8080",
		//		ErrorLog: App.errorLog,
		Handler: App.routes(),
	}
	go HandleWeb()
	cfmt.PrintlnOk("Web server initialized")
}

func HandleWeb() {
	err := App.srv.ListenAndServe()
	log.Fatal(err)
}
