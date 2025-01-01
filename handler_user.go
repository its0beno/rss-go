package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/its0benp/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type paramters struct{
		Name string `json:"name"`
	}
	decoder :=json.NewDecoder(r.Body)
	params:=paramters{}
	err := decoder.Decode(&params)
	if err!=nil {
		responseWithError(w,400,fmt.Sprintf("Error parsing json %v",err))
		return
	}
	user,err := apiCfg.DB.CreateUser(
		r.Context(), database.CreateUserParams{
			Name: params.Name,
			ID: uuid.New(),
			CreatedAt:time.Now().UTC() ,
			UpdateAt: time.Now().UTC(),
		}	)
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("error creating user %v ",err))
		return
	}
	update_user := databaseUserToUser(user)
	responseWithJson(w,201,update_user)
}
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	
	update_user := databaseUserToUser(user)

	responseWithJson(w, 200, update_user)
	}	