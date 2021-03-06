package privat

import (
	"encoding/xml"
	"bytes"
	"log"
	"time"
)

type BalanceRequestXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data struct {
		XMLName xml.Name `xml:"data"`
		Oper    string   `xml:"oper"`
		Wait    int      `xml:"wait"`
		Test    int      `xml:"test"`
		Payment Payment
	}
}

type BalanceResponseXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  float64  `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data struct {
		XMLName xml.Name `xml:"data"`
		Oper    string   `xml:"oper"`
		Info struct {
			XMLName     xml.Name `xml:"info"`
			Cardbalance Cardbalance
		}
	}
}

type Cardbalance struct {
	XMLName    xml.Name  `xml:"cardbalance"`
	Card       Card
	AvBalance  float64   `xml:"av_balance"`
	BalDate    time.Time `xml:"bal_date"`
	BalDyn     string    `xml:"bal_dyn"`
	Balance    float64   `xml:"balance"`
	FinLimit   float64   `xml:"fin_limit"`
	TradeLimit float64   `xml:"trade_limit"`
}

type Card struct {
	XMLName        xml.Name `xml:"card"`
	Account        int      `xml:"account"`
	CardNumber     int      `xml:"card_number"`
	AccName        string   `xml:"acc_name"`
	AccType        string   `xml:"acc_type"`
	Currency       string   `xml:"currency"`
	CardType       string   `xml:"card_type"`
	MainCardNumber int      `xml:"main_card_number"`
	CardStat       string   `xml:"card_stat"`
	Src            string   `xml:"src"`
}

func (api *Privat24Api) GetBalance(cardNumber string) BalanceResponseXML {
	url := api.apiUrl + "/balance"

	payment:=new(Payment)
	payment.SetBalanceProperties(cardNumber, "UA")

	balanceRequest := new(BalanceRequestXML)
	balanceRequest.Version = "1.0"
	balanceRequest.Data.Oper = "cmt"
	balanceRequest.Data.Wait = 0
	balanceRequest.Data.Test = 0
	balanceRequest.Data.Payment=*payment
	balanceRequest.Merchant=api.getMerchantStruct(balanceRequest.Data)

	bytexml, err := xml.Marshal(balanceRequest)
	if err != nil {
		log.Println(err.Error())
	}

	reqBody := append([]byte{}, []byte(xml.Header)...)
	reqBody = append(reqBody, bytexml...)

	response, err := api.requestXML(url, bytes.NewBuffer(bytexml), "POST")
	if err != nil {
		log.Println(err.Error())
	}

	balanceResponse := new(BalanceResponseXML)

	err = xml.Unmarshal(response, &balanceResponse)
	if err != nil {
		log.Println(err.Error())
	}

	return *balanceResponse
}
