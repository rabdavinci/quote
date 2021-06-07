package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var qs Quotes

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(qs)
}

func getQuoteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}

	if id == 0 {
		panic("Id required")
	}

	i := qs.FindIndexById(id)

	json.NewEncoder(w).Encode(qs[i])
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	ri := rand.Intn(len(qs))
	json.NewEncoder(w).Encode(qs[ri])
}

func getQuotesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	json.NewEncoder(w).Encode(qs.GetAllByCategory(name))
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(qs.Create(reqBody))
}

func updateQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	if id == 0 {
		panic("Id required")
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	i := qs.FindIndexById(id)
	json.NewEncoder(w).Encode(qs.Update(reqBody, i))
}

func deleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	if id == 0 {
		panic("Id required")
	}
	i := qs.FindIndexById(id)
	qs.DeleteByIndex(i)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getAllQuotes).Methods("GET")
	router.HandleFunc("/quote/{id}", getQuoteById).Methods("GET")
	router.HandleFunc("/random-quote", getRandomQuote)
	router.HandleFunc("/category/{name}", getQuotesByCategory)
	router.HandleFunc("/quote", createQuote).Methods("POST")
	router.HandleFunc("/quote/{id}", updateQuote).Methods("PUT")
	router.HandleFunc("/quote/{id}", deleteQuote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	qs.FillWithTestData(5)
	go garbageWorker()
	handleRequests()
}
