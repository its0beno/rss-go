package main

import (
	"fmt"
	"net/http"

	"github.com/its0benp/rssagg/internal/auth"
	"github.com/its0benp/rssagg/internal/database"
)

type authHandeler func(http.ResponseWriter, *http.Request, database.User)


func (cfg *apiConfig) middlwareAuth(handler authHandeler) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		apiKey,err := auth.GetAPiKey(r.Header)
		if err != nil{
			responseWithError(w, 403,fmt.Sprintf("Auth error %v",err))
			return
		}
		user,err :=cfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err != nil{
			responseWithError(w, 404,fmt.Sprintf("Couldnt Get User %v",err))
			return
		}
		handler(w,r,user)
	}
}