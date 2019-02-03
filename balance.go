package privat

import (
	"encoding/xml"
	"bytes"
	"log"
	"time"
)

type Cardbalance struct {
	XMLName     xml.Name  `xml:"cardbalance"`
	Card        Card
	av_balance  float64   `xml:"av_balance"`
	bal_date    time.Time `xml:"bal_date"`
	bal_dyn     string    `xml:"bal_dyn"`
	balance     float64   `xml:"balance"`
	fin_limit   float64   `xml:"fin_limit"`
	trade_limit float64   `xml:"trade_limit"`
}
type Card struct {
	XMLName          xml.Name `xml:"card"`
	account          int      `xml:"account"`
	card_number      int      `xml:"card_number"`
	acc_name         string   `xml:"acc_name"`
	acc_type         string   `xml:"acc_type"`
	currency         string   `xml:"currency"`
	card_type        string   `xml:"card_type"`
	main_card_number int      `xml:"main_card_number"`
	card_stat        string   `xml:"card_stat"`
	src              string   `xml:"src"`
}

func (api *Privat24Api) GetBalance(cardNumber string) {
	url := api.apiUrl + "/balance"

	data := new(Data)
	items := make([]interface{}, 4)

	paymentProp := make([]Prop, 2)
	paymentProp = append(paymentProp, Prop{Name: "cardnum", Value: cardNumber}, Prop{Name: "country", Value: "UA"})

	payment := new(Payment)
	payment.Properties = paymentProp

	items = append(items, new(Oper), new(Wait), new(Test), payment)

	data.Items = items

	queryXML := new(QueryXML)
	queryXML.setMerchant(api.merchantID, api.merchantPassword)
	queryXML.Data = *data

	xml, err := xml.Marshal(queryXML)
	if err != nil {
		log.Println(err.Error())
	}

	res, err := api.requestXML(url, bytes.NewBuffer(xml))
	if err != nil {
		log.Println(err.Error())
	}

}
