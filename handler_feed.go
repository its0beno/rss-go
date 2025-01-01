package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/its0benp/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	type paramters struct{
		Name string `json:"name"`
		Url string `json:"url"`

	}
	decoder :=json.NewDecoder(r.Body)
	params:=paramters{}
	err := decoder.Decode(&params)
	if err!=nil {
		responseWithError(w,400,fmt.Sprintf("Error parsing json %v",err))
		return
	}
	feed ,err := apiCfg.DB.Createfeed(
		r.Context(), database.CreatefeedParams{
			Name: params.Name,
			ID: uuid.New(),
			CreatedAt:time.Now().UTC() ,
			UpdateAt: time.Now().UTC(),
			Url: params.Url,
			UserID: user.ID,
		}	)
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("error creating user %v ",err))
		return
	}

	responseWithJson(w,201,databaseFeedToFeed(feed))
}
func (apiCfg *apiConfig) handlerGeteFeed(w http.ResponseWriter, r *http.Request, ){
	
	feed ,err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("Couldnt get feeds %v ",err))
		return
	}

	responseWithJson(w,201,databaseFeedsToFeeds(feed))
}
