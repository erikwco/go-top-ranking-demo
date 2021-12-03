package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"top-ranking/internal/cli"
	"top-ranking/internal/parser"
)

func GetTopRanked(w http.ResponseWriter, r *http.Request) {

	//*************************************************
	// Get Query Parameters
	//*************************************************
	top, err := strconv.Atoi(r.URL.Query().Get("top"))
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("{message: 'Error on top parameter [%v]', code: -1 }", err)))
	}
	lang := r.URL.Query().Get("language")
	if top < 0 {
		top = 1
	}

	//*************************************************
	// Get CSV Data
	//*************************************************
	records, err := cli.GetCSVData()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("{message: 'Error getting CSV information [%v]', code: -1 }", err)))
	}

	//*************************************************
	// Filtering by Language
	//*************************************************
	filtered, l := parser.FilterByLanguage(records, lang)
	if top > l {
		top = l
	}

	//*************************************************
	// Returning
	//*************************************************
	if l > 0 {
		// Convirtiendo a Json y devolviendo el top solicitado
		d, err := parser.ConverToJson(filtered[:top])
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("{message: 'Error parsing data [%v]', code: -1 }", err)))
		}
		w.WriteHeader(200)
		w.Write(d)
	} else {
		w.WriteHeader(404)

		w.Write([]byte("{message: 'Language not found', code: -1 }"))
	}

}

func main() {

	// handler
	http.HandleFunc("/top", GetTopRanked)

	// simple server
	fmt.Println("Starting server on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error starting server %v", err)
	}

}
