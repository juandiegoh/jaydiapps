package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/juandiegoh/jaydiapps/love"
)

// Index controller of /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to JayDiApps!\n")
}

// LoveIndex controller of /loves
func LoveIndex(w http.ResponseWriter, r *http.Request) {
	responseAsJSON(w, love.FindAll(), http.StatusOK)
}

// LoveShow return Love with id {id} or 404
func LoveShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var loveID int
	var err error
	if loveID, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}
	l := love.FindByID(int32(loveID))
	if l.ID > 0 {
		responseAsJSON(w, l, http.StatusOK)
		return
	}

	// If we didn't find it, 404
	responseAsJSON(w, jsonErr{Code: http.StatusNotFound, Text: "Not Found"}, http.StatusNotFound)
}

// LoveRandom return a random value from Loves
func LoveRandom(w http.ResponseWriter, rq *http.Request) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ls := love.FindAll()
	l := ls[r.Intn(len(ls))]
	responseAsJSON(w, l, http.StatusOK)
}

func responseAsJSON(w http.ResponseWriter, o interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(o); err != nil {
		panic(err)
	}
}
