package main

import (
	"encoding/json"
	"fmt"
	"goroutines/authentication/jwt/Token"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	fmt.Println("Starting service!")
	// mux:=mux.NewRouter()
	// mux.HandleFunc("/jwt", middleware(loginHandler)).Methods("POST")
	// mux.HandleFunc("/user", middleware(userdetails)).Methods("POST")
	// http.ListenAndServe("localhost:3000", mux)

	password := "password"
	pwd,_:=bcrypt.GenerateFromPassword([]byte(password),14)
	fmt.Println(string(pwd))

	err:=bcrypt.CompareHashAndPassword(pwd,[]byte(password))
	fmt.Println(err==nil)

}

type User struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Came to middleware:", r.URL.Path, r.URL.User)
		h(w, r)
	}
}
func loginHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error %s:", err.Error())))
		return
	}

	if u.UserName == "admin" && u.Password == "12345" {
		token, err := token.CreateToken(u.UserName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("unable to generate token with error %s:", err.Error())))
			return
		}
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, fmt.Sprintf("Invalid Credentials"))
	}
}
func userdetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		w.Write([]byte("Authorization required"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	t:=tokenString[len("Bearer "):]
	err:=token.VerifyToken(t)
	if err!=nil{
		w.Write([]byte("Not a valid Token"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Valid User"))
	w.WriteHeader(http.StatusOK)
}
