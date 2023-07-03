package server

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"

	"github.com/shopspring/decimal"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Name    string   `xml:"name,attr"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string          `xml:"ID,attr"`
	NumCode  string          `xml:"NumCode"`
	CharCode string          `xml:"CharCode"`
	Nominal  int64           `xml:"Nominal"`
	Name     string          `xml:"Name"`
	Value    decimal.Decimal `xml:"Value"`
}

func replace(str string) string {
	return strings.ReplaceAll(str, ",", ".")
}

func (vc *ValCurs) Parse(xmlCourses string) error {
	data := replace(xmlCourses)
	err := xml.Unmarshal([]byte(data), vc)
	if err != nil {
		return err
	}
	return nil
}

// получение XML документа со списком валют
func GetData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	sb := string(body)
	return sb, nil
}
