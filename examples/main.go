package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

var dvs []Device
var version string

func init() {
	version = "2.10.9"
	dvs = []Device{
		{1, "13414214242424", version},
		{2, "1", version},
		{3, "3", version},
	}
}

func main() {
	dMux := http.NewServeMux()
	dMux.HandleFunc("/devices", getDevices)
	log.Fatal(http.ListenAndServe(":8088", dMux))
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
