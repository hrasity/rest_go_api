package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type creddentials struct {
		Username string `json:"email"`
		Password string `json:"password"`
	}

	var creds creddentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)

	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "Invalid Json supplied, or json entirely missing"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	app.infoLog.Println(creds.Username, creds.Password)
	payload.Error = false
	payload.Message = "Signed in"

	err = app.writeJSON(w, http.StatusOK, payload)

	if err != nil {
		app.errorLog.Println(err)
	}

}
