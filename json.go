package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func responseWithError(w http.ResponseWriter, code int, msg string){
	if code >499{
		log.Println("Responding WIth 5xx error:", msg)
	}
	type errResponse struct{
		Error string "json:error"
	}
	responseWithJson(w,code,errResponse{
		Error: msg,
	})
}

func responseWithJson(w http.ResponseWriter,code int , payload interface{}){

	dat, err :=json.Marshal(payload)
	if err != nil {
		log.Printf("Failed Marshal Json Response: %v ",payload)
		w.WriteHeader(500)
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}