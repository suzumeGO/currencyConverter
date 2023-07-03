package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

type Currencies struct {
	FirstCurrency  string          `json:"firstCurrency"`
	Amount         decimal.Decimal `json:"amount"`
	SecondCurrency string          `json:"secondCurrency"`
}

func Desirialize(rawMsg *http.Request) (Currencies, error) {
	decoder := json.NewDecoder(rawMsg.Body)
	curr := Currencies{}
	err := decoder.Decode(&curr)
	if err != nil {
		return Currencies{}, fmt.Errorf("unmarshal: %w", err)
	}
	return curr, nil
}
