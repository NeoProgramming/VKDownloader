package core

import (
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"vkdownloader/cfmt"
)

// The serverError helper writes an error message to the errorLog
// and then sends a 500 "Internal Server Error" response to the user.
func (app *Application) serverError(w http.ResponseWriter, err error) {
	cfmt.PrintlnErr("serverError: ", err.Error(), debug.Stack())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description to user.
func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// It's just a convenience wrapper around clientError that sends a "404 Page Not Found" response to the user.
func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func GetGlobalIP() string {
	getip := "https://api.country.is"
	resp, err := http.Get(getip)
	if err != nil {
		cfmt.PrintlnErr("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cfmt.PrintlnErr("Error:", err)
		return ""
	}
	return string(body)
}
