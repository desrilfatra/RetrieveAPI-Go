package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"retrieveapi-go/entity"
	"time"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Data", getData)
	fmt.Println("Now loading on port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func getData(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var users []entity.User
	// json.Unmarshal(body, &users)
	// fmt.Println(users)
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		log.Fatal(err)
	}

	jsonData, _ := json.Marshal(&users)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
