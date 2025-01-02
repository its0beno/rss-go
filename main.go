package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/its0benp/rssagg/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}
func main () {
	feed,err:= urlTOFeedFollows("https://www.wagslane.dev/index.xml")
	fmt.Println(feed,err)
	godotenv.Load()
	portString:=os.Getenv("PORT")
	if portString=="" {
	 log.Fatal("PORT environment variable was not set")
	}
	fmt.Print(portString)

	DB_URL:=os.Getenv("DB_URL")
	if DB_URL=="" {
	 log.Fatal("DB url was not found ")
	}
	conn,err :=sql.Open("postgres",DB_URL)
	if err !=nil{
		log.Fatal("Cant connect to database",err)
	}
	queries:=database.New(conn)
	
	apiCfg :=apiConfig{
		DB:queries,
	}
	router:= chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: [] string{"https://*","http://*"},
		AllowedMethods: [] string{"GET","OPTIONS", "PUT","POST","DELETE"},
		AllowedHeaders: [] string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	V1router:= chi.NewRouter()
	V1router.Get("/healthz", handlerReadiness )
	V1router.Get("/err",handlerErr)
	V1router.Post("/users",apiCfg.handlerCreateUser)
	V1router.Get("/users",apiCfg.middlwareAuth(apiCfg.handlerGetUser))
	V1router.Post("/feeds",apiCfg.middlwareAuth(apiCfg.handlerCreateFeed))
	V1router.Post("/follow/feed", apiCfg.middlwareAuth(apiCfg.handlerCreateFeedFollow))
	V1router.Get("/follow/feed", apiCfg.middlwareAuth(apiCfg.handlerGeteFeedFollow))
	V1router.Delete("/follow/feed/{feed_id}", apiCfg.middlwareAuth(apiCfg.handlerDeleteFeedFollow))
	V1router.Get("/feeds" ,apiCfg.handlerGeteFeed)
	router.Mount("/v1",V1router)
	srv := &http.Server{Handler:router,Addr:":"+portString}
	log.Printf("serverRunig On Port: %v", portString)
	err = srv.ListenAndServe()
	if err!=nil {
		log.Fatal(err)
	}
}