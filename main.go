package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func main(){
	r := mux.NewRouter()

	var domain string = "localhost"

	r.Host(domain).
		Path("/").
		HandlerFunc(HomeHandler).
		Name("root")

	r.Host("{subdomain}."+domain).
		Path("/").
		HandlerFunc(NewsHandler).
		Name("subRoot")
	newsHost, err := r.Get("subRoot").URLHost("subdomain", "news")
	chk(err)

	r.Host("{subdomain}."+domain).
		Path("/").
		HandlerFunc(ApiHandler).
		Name("api")
	apiHost, err := r.Get("api").URLHost("subdomain", "api")
	chk(err)

	log.Println(apiHost)
	log.Println(newsHost)
	log.Fatal(http.ListenAndServe(":80", r))
}
func ApiHandler(w http.ResponseWriter, r *http.Request){

	log.Println("home")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	log.Println(vars)
	fmt.Fprintf(w, "hello from news")
}
func NewsHandler(w http.ResponseWriter, r *http.Request){

	log.Println("home")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	log.Println(vars)
	fmt.Fprintf(w, "hello from news")
}
func HomeHandler(w http.ResponseWriter, r *http.Request){

	log.Println("home")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	log.Println(vars)
	fmt.Fprintf(w, "hello World")
}
func chk(err error){
	if err != nil{
		panic(err)
	}

}