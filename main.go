package main

import (
	conv "Projects/currencyConverter/server"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const (
	portNumber               = ":8081"
	notAllowMethodStatusCode = 405
	inputErr                 = "Input data error"
	allowHeader              = "Allow"
	notAllowMethod           = "Method not allowed!!"
	dateFormat               = "2006-01-02 15:04:05"
	serverStartString        = "Web-server starting on http://localhost"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set(allowHeader, http.MethodPost)
		http.Error(w, notAllowMethod, notAllowMethodStatusCode)
		return
	}
	reqCurr, err := conv.Desirialize(r)
	if err != nil {
		http.Error(w, inputErr, notAllowMethodStatusCode)
		return
	}
	value, err := conv.GetExchangeRate(reqCurr.FirstCurrency, reqCurr.Amount, reqCurr.SecondCurrency)
	if err != nil {
		http.Error(w, inputErr, notAllowMethodStatusCode)
		return
	}
	response := struct {
		MsgTime string
		Value   decimal.Decimal
	}{time.Now().Format(dateFormat), value}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Println(serverStartString + portNumber)
	err := http.ListenAndServe(portNumber, mux)
	log.Fatal(err)
}
