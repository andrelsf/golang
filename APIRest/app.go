package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
)

const (
	APP_PORT = ":8000"
)

type (
	ResponseStatus struct {
		Status		int
		Message		string
	}
	/*
		Entidade User
		`json:"-"` Oculta campo na resposta do JSON
	*/
	User struct {
		Name		string		`json:"name"`
		Username	string		`json:"username"`
		Email		string		`json:"email"`
		Password	string		`json:"-"`
		Active		bool		`json:"active"`
		CreatedAt	time.Time	`json:"createat"`
		UpdatedAt	time.Time	`json:"updateat"`
	}
)

/*
	POST: /api/registry
	@CURL: 
	curl -X POST -H "Content-Type: application/json" \
	-d '{"Name":"Andre","Username":"andre.ferreira","Email":"andre.ferreira@gmail.com","Password":"asdf12j3h4k1j"}' \
	http://localhost:8000/api/registry
*/
func Registry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		user.Active = false
		user.CreatedAt = time.Now().Local()

		userJson, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
	} else {
		ResponseResourceNotImplemented(w)
	}
}

/*
	TODO
	GET: /api/users - Get all Users
	@CURL: curl -X GET -H "Content-type: application/json" localhost:8000/api/users
*/
func Users(w http.ResponseWriter, r *http.Request) {
	ResponseResourceNotImplemented(w)
}

/*
	GET: /api/ping
	@CURL: curl -X GET -H "Content-type: application/json" localhost:8000/api/ping
*/
func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		respStatus := ResponseStatus{Status: 200, Message: "pong",}

		respJson, err := json.MarshalIndent(respStatus, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	} else {
		ResponseResourceNotImplemented(w)
	}
}

func ResponseResourceNotImplemented(w http.ResponseWriter) {
	resp := ResponseStatus{Status: http.StatusNotImplemented, Message: "Resource not implemented"}
		
	respJson, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write(respJson)
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
	mux.Handle("/api/ping", loggerMiddleware(http.HandlerFunc(Ping)))
	mux.Handle("/api/users", loggerMiddleware(http.HandlerFunc(Users)))
	mux.Handle("/api/registry", loggerMiddleware(http.HandlerFunc(Registry)))
	fmt.Println("Server is running on:", APP_PORT)
	http.ListenAndServe(APP_PORT, mux)
}