package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type PlacementReq struct {
	Name      string `json:"name"`
	RequestID string `json:"requestId"`
}

type PlacementRes struct {
	ReqHandled bool   `json:"reqHandled"`
	RequestID  string `json:"requestId"`
}

var reqPool sync.Pool
var resPool sync.Pool

func placementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	preq, ok := reqPool.Get().(*PlacementReq)
	if !ok {
		preq = &PlacementReq{}
	}
	defer reqPool.Put(preq)
	err := json.NewDecoder(r.Body).Decode(preq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// do something randomly for (3 + r)ms
	duration := time.Duration(rand.Intn(10)+3) * time.Millisecond
	time.Sleep(duration)
	log.Printf("Serving: %s\n", preq.RequestID)
	pres, ok := resPool.Get().(*PlacementRes)
	if !ok {
		pres = &PlacementRes{}
	}
	defer resPool.Put(pres)
	pres.ReqHandled = true
	pres.RequestID = preq.RequestID
	err = json.NewEncoder(w).Encode(pres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/placement", placementHandler)
	log.Println("Serving '/placement' on :1771")
	log.Fatal(http.ListenAndServe(":1771", nil))
}
