package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, "index", "Test")
}

/* UTILITY FUNCATIONS */
func Json(object interface{}) []byte {
	j, err := json.Marshal(object)
	if err != nil {
		fmt.Println("error:", err)
	}
	return j
}
