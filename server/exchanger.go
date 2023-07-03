package server

import (
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

const (
	url        = "https://www.cbr-xml-daily.ru/daily_utf8.xml"
	recieveErr = "Recieve data error"
)

var errNotContains error = errors.New("no such currency")

type rate map[string]decimal.Decimal

func GetValutesRates() rate {
	valutes := make(map[string]decimal.Decimal, 30)

	data, err := GetData(url)
	if err != nil {
		log.Println(recieveErr)
	}
	var res ValCurs
	res.Parse(data)

	for _, val := range res.Valutes {
		valutes[val.CharCode] = val.Value.Div(decimal.NewFromInt(val.Nominal))
		valutes["RUB"] = decimal.NewFromInt(1)
	}
	return valutes
}

func GetExchangeRate(srcVal string, amount decimal.Decimal, dstVal string) (decimal.Decimal, error) {
	val, ok := GetValutesRates()[srcVal]
	if !ok {
		return val, errNotContains
	}
	val, ok = GetValutesRates()[dstVal]
	if !ok {
		return val, errNotContains
	}
	return GetValutesRates()[srcVal].Mul(amount).DivRound(GetValutesRates()[dstVal], 4), nil
}
