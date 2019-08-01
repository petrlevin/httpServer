package main

import (
	"net/http"
	 "simpleServer/srv"
)


func main(){
	srv := http.Server{
		Addr:":8080",
		Handler:&srv.Handler{

		},
	}
	srv.ListenAndServe()
}