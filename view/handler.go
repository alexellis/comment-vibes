package function

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.Get("http://gateway.openfaas:8080/function/comments")

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	results := []Result{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, results)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(tpl.Bytes())
}

type Result struct {
	Emoji string `json:"emoji"`
	Total int    `json:"total"`
}
