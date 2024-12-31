package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)
func main () {

	godotenv.Load()
	portString:=os.Getenv("PORT")
	if portString=="" {
	 log.Fatal("PORT environment variable was not set")
	}
	fmt.Print(portString)
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
	router.Mount("/v1",V1router)
	srv := &http.Server{Handler:router,Addr:":"+portString}
	log.Printf("serverRunig On Port: %v", portString)
	err := srv.ListenAndServe()
	if err!=nil {
		log.Fatal(err)
	}
}