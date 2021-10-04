package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Time struct {
	Time string `json:"time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	data := &Time{Time: t}
	res, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(res)

}

func main() {
	http.HandleFunc("/time", handler)
	http.ListenAndServe(":5000", nil)
}
