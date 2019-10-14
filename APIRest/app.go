package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

const APP_PORT = ":8000"

type ResponseStatus struct {
	Status		int
	Message		string
}

/*
	GET: /api/ping
	@CURL: curl -X GET -H "Content-type: application/json" localhost:8000
*/
func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		respStatus := ResponseStatus{Status: 200, Message: "pong",}

		respJson, err := json.MarshalIndent(respStatus, "", "  ")
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	} else {
		resp := ResponseStatus{Status: http.StatusNotImplemented, Message: "Resource not implemented"}
		
		respJson, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write(respJson)
	}
}

/*
	MiddleWare para fazer log de todas as requisições
*/
func loggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Request received: %v\n", r)
		nextHandler.ServeHTTP(w, r)
		fmt.Println("Request handled successfully")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/ping", loggerMiddleware(http.HandlerFunc(ping)))
	fmt.Println("Server is running on:", APP_PORT)
	http.ListenAndServe(APP_PORT, mux)
}