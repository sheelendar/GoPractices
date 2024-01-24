package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var muxRouter *mux.Router

func init() {
	muxRouter = mux.NewRouter
}
func StartHTTPServer() {
	muxRouter.HandleFunc("hello", GetHelloHandler).Methods("GET")
	http.Handle("/", muxRouter)
	http.ListenAndServe(":8080", nil)
}
func GetHelloHandler(res http.ResponseWriter, req *http.Request) {
	//req.URL.Query().Get("param")

	jonsRes, _ := json.Marshal("hello")
	res.Write(jonsRes)
}
