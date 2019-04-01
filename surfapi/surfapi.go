package main

import (
	"encoding/json"
	"net/http"

	"./datautil"
)

func main() {
	http.HandleFunc("/", dataToJson)
	http.ListenAndServe(":3001", nil)
}

func dataToJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")

	url := "https://www.ndbc.noaa.gov/data/realtime2/46029.spec"
	rawBouyData := datautil.GetBouyData(url)
	bouyData := datautil.HandleRawData(rawBouyData)
	packagedBouyData := datautil.DataToStructs(bouyData)

	js, err := json.Marshal(packagedBouyData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
