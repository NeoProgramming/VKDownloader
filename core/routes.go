package core

import (
	"net/http"
	"vkdownloader/cfmt"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/images/favicon.ico")
}

func (app *Application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/favicon.ico", faviconHandler)

	mux.HandleFunc("/", app.home)

	// POST handlers
	mux.HandleFunc("/set-app-id", app.setAppId)
	mux.HandleFunc("/set-app-url", app.setAppToken)
	mux.HandleFunc("/download-album", app.downloadAlbum)
	mux.HandleFunc("/download-owner", app.downloadOwner)
	mux.HandleFunc("/stop-worker", app.stopWorker)
	mux.HandleFunc("/get-ip", app.getIP)
	mux.HandleFunc("/exit", app.exit)

	cfmt.PrintlnOk("Routes initialized")
	return mux
}
