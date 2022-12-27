package main

import (
	"net/http"
)

// healthcheckHandler is implemented as a
// method on our application struct		!!!!
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	// passing data to json.Marshal() func
	// and showing data
	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}