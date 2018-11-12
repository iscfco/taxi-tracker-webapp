package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"os"
	"taxi-tracker-webapp/constants"
	"taxi-tracker-webapp/model"
	"taxi-tracker-webapp/utils"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")

	c, err := r.Cookie("session")
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error interno"))
		return
	}

	if c != nil {
		var session model.Session
		json.Unmarshal([]byte(c.Value), &session)
		statusCode, jsonBody, err := utils.DoGet("", constants.UrlGetService, session.AccessToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error interno"))
			return
		}
		var taxiService model.TaxiService
		err = json.Unmarshal([]byte(jsonBody), &taxiService)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error interno"))
			return
		}
		var nextPage string
		if statusCode == 200 {
			buf := new(bytes.Buffer)
			if taxiService.CustomerId != "" {
				nextPage = os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/view/dashboard/track_service.html"
				tmpl, err := template.ParseFiles(nextPage)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error interno"))
					return
				}
				if err = tmpl.Execute(buf, ""); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error interno"))
					return
				}
			} else {
				nextPage = os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/view/dashboard/get_service.html"
				tmpl, err := template.ParseFiles(nextPage)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error interno"))
					return
				}
				if err = tmpl.Execute(buf, ""); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error interno"))
					return
				}
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(buf.String()))
			return
		}
		return
	}

	indexTemplate := os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/view/signin.html"
	tmpl, err := template.ParseFiles(indexTemplate)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, ""); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(buf.String()))
	return
}
