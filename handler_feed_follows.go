package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/its0benp/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type paramters struct{
		FeedID uuid.UUID `json:"feed_id"`


	}
	decoder :=json.NewDecoder(r.Body)
	params:=paramters{}
	err := decoder.Decode(&params)
	if err!=nil {
		responseWithError(w,400,fmt.Sprintf("Error parsing json %v",err))
		return
	}
	feed_follow ,err := apiCfg.DB.CreatefeedFollow(
		r.Context(), database.CreatefeedFollowParams{
			ID: uuid.New(),
			CreatedAt:time.Now().UTC() ,
			UpdateAt: time.Now().UTC(),
			FeedID: params.FeedID,
			UserID: user.ID,
		}	)
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("error creating user %v ",err))
		return
	}

	responseWithJson(w,201,databaseFeedFollowToFeedFollow(feed_follow))
}
func (apiCfg *apiConfig) handlerGeteFeedFollow(w http.ResponseWriter, r *http.Request,user database.User ){
	
	feed_follow ,err := apiCfg.DB.GetUsersFeeds(r.Context(),user.ID)
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("Couldnt get feeds %v ",err))
		return
	}

	responseWithJson(w,201,databaseFeedFollowsToFeedFollows(feed_follow))
}
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request,user database.User ){
	FeedFollowIdStr := chi.URLParam(r,"feed_id")
	feedFollowUid, err := uuid.Parse(FeedFollowIdStr)


	if err!=nil {
		responseWithError(w,400,fmt.Sprintf("Couldnt parse Follow feed id  %v",err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: feedFollowUid,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("Couldnt delete feed %v ",err))
		return
	}

	responseWithJson(w,200,struct{}{})
}
