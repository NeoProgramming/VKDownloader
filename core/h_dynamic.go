package core

import (
	"fmt"
	"net/http"
	"vkdownloader/cfmt"
)

func (app *Application) setAppId(w http.ResponseWriter, r *http.Request) {
	cfmt.PrintlnFunc("setAppId")
	if r.Method == http.MethodPost {
		app.config.AppID = r.FormValue("app_id")
		cfmt.PrintlnLine("app_id:", app.config.AppID)
		// ... update user information in the database ...
		w.Write([]byte("setAppId ok"))
		SaveConfig()
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (app *Application) setAppToken(w http.ResponseWriter, r *http.Request) {
	cfmt.PrintlnFunc("setAppToken")
	if r.Method == http.MethodPost {
		urlStr := r.FormValue("app_url")
		if urlStr != "" {
			app.config.AccessToken = extractAccessToken(urlStr)
			app.config.RecentIP = GetGlobalIP()
		}
		cfmt.PrintlnLine("app_token:", app.config.AccessToken)
		// ... update user information in the database ...
		w.Write([]byte("setAppToken ok"))
		SaveConfig()
		InitVK()
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (app *Application) downloadAlbum(w http.ResponseWriter, r *http.Request) {
	if app.vk == nil {
		InitVK()
	}
	if !app.running {
		app.running = true
		app.wg.Add(1)
		album := r.FormValue("id")
		go app.worker(album)

		fmt.Fprint(w, "true")
		cfmt.PrintlnImp("Worker started: ", album)
	}
}

func (app *Application) downloadOwner(w http.ResponseWriter, r *http.Request) {
	if app.vk == nil {
		InitVK()
	}
	if !app.running {
		app.running = true
		app.wg.Add(1)
		owner := r.FormValue("id")
		go app.worker(owner)

		fmt.Fprint(w, "true")
		cfmt.PrintlnImp("Worker started: ", album)
	}
}


func (app *Application) stopWorker(w http.ResponseWriter, r *http.Request) {
	if app.running {
		app.running = false
		app.wg.Wait()
		fmt.Fprint(w, "false")
		cfmt.PrintlnImp("Worker stopped manually")
	}
}

func (app *Application) getIP(w http.ResponseWriter, r *http.Request) {
	app.config.RecentIP = GetGlobalIP()
	cfmt.PrintlnFunc("getIP: ", app.config.RecentIP)
	fmt.Fprint(w, app.config.RecentIP)
}

// stop handler
func (app *Application) exit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(200)
		fmt.Fprint(w, "Shutting down the server...")
		cfmt.PrintlnImp("Shutting down the server...")
		go func() {
			cfmt.PrintlnImp("Server stopped.")
		}()
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
