package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"taxi-tracker-webapp/constants"
	"taxi-tracker-webapp/model"
	"taxi-tracker-webapp/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	r.ParseForm()
	user := model.User{
		User:     r.Form["user"][0],
		Password: r.Form["password"][0],
	}

	statusCode, jsonBody, err := utils.DoPost(user, constants.UrlLogin, "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error interno"))
		return
	}
	var customerSession model.CustomerSession
	err = json.Unmarshal([]byte(jsonBody), &customerSession)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error interno"))
		return
	}

	if statusCode == 200 && customerSession.Res.ResCode == "0" {
		session := model.Session{
			AccessToken:  customerSession.AccessToken,
			RefreshToken: customerSession.RefreshToken,
		}
		sessionStr, _ := json.Marshal(session)
		cookie := http.Cookie{
			Name:  "session",
			Value: string(sessionStr),
		}
		http.SetCookie(w, &cookie)

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
	}

	if statusCode == 200 && customerSession.Res.ResCode == "EUS001" {
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

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Error interno"))
	return
}

func postLogin(user, pwd string) {
	jsonData := map[string]string{"user": user, "password": pwd}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", "http://localhost:5000/api/customer_session/", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		fmt.Println(response)
	}
}
