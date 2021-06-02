package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Quote struct {
	Id         int    `json:"Id"`
	Author     string `json:"Author"`
	Quote      string `json:"Quote"`
	Category   string `json:"Category"`
	Created_at int64  `json:"Created_at"`
}

var Quotes []Quote

func allQuotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: allQuotes")
	json.NewEncoder(w).Encode(Quotes)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, quote := range Quotes {
		if quote.Id == key {
			json.NewEncoder(w).Encode(quote)
		}
	}
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	RandomIndex := rand.Intn(len(Quotes))
	json.NewEncoder(w).Encode(Quotes[RandomIndex])
}

func getQuotesByCategory(w http.ResponseWriter, r *http.Request) {
	var FilteredQuotes []Quote
	vars := mux.Vars(r)
	category := vars["id"]
	for _, quote := range Quotes {
		if quote.Category == category {
			FilteredQuotes = append(FilteredQuotes, quote)
		}
	}
	json.NewEncoder(w).Encode(FilteredQuotes)
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var quote Quote
	quote.Id = getMaxQuoteId() + 1
	json.Unmarshal(reqBody, &quote)
	Quotes = append(Quotes, quote)

	json.NewEncoder(w).Encode(quote)
}

func updateQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, quote := range Quotes {
		if quote.Id == id {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var quote Quote
			quote.Id = getMaxQuoteId() + 1
			json.Unmarshal(reqBody, &quote)
			Quotes[index] = quote

			json.NewEncoder(w).Encode(quote)
		}
	}
}

func deleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, quote := range Quotes {
		if quote.Id == id {
			Quotes = append(Quotes[:index], Quotes[index+1:]...)
		}
	}
}

func getMaxQuoteId() (r int) {
	for _, quote := range Quotes {
		if quote.Id > r {
			r = quote.Id
		}
	}
	return
}

func addTestQuotes(count int) {
	for i := 1; i <= count; i++ {
		var quote Quote
		quote.Id = getMaxQuoteId() + 1
		quote.Author = fmt.Sprintf("Author %d", i)
		quote.Quote = fmt.Sprintf("Quote %d", i)
		quote.Category = fmt.Sprintf("Category %d", i)
		quote.Created_at = time.Now().Unix()
		Quotes = append(Quotes, quote)
	}
}

func garbageWorker() {
	ti := 1 * time.Hour
	tl := time.Now().Add(-ti).Unix()
	removed := 0
	for i := 0; i < len(Quotes); i++ {
		if Quotes[i].Created_at < tl {
			Quotes = append(Quotes[:i], Quotes[i+1:]...)
			removed++
			i--
		}
	}
	time.Sleep(ti)
	go garbageWorker()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", allQuotes)
	router.HandleFunc("/quote/{id}", getQuote)
	router.HandleFunc("/random-quote", getRandomQuote)
	router.HandleFunc("/category/{id}", getQuotesByCategory)
	router.HandleFunc("/quote", createQuote).Methods("POST")
	router.HandleFunc("/quote/{id}", updateQuote).Methods("PUT")
	router.HandleFunc("/quote/{id}", deleteQuote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	addTestQuotes(5)

	go garbageWorker()
	handleRequests()
}
