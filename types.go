package privat

import (
	"encoding/xml"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	//"log"
)

type Merchant struct {
	ID        int    `xml:"id"`
	Signature string `xml:"signature"`
}

type Prop struct {
	XMLName xml.Name `xml:"prop"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Payment struct {
	XMLName    xml.Name `xml:"payment"`
	ID         string   `xml:"id,attr"`
	Properties []Prop
}

type Info struct {
	XMLName xml.Name `xml:"info"`
	Items   []interface{}
}

// установить prop для просмотра баланса
// cardnum Номер карты
// country Страна
func (p *Payment) SetBalanceProperties(cardnum string, country string) {
	paymentProp := make([]Prop, 2)
	paymentProp[0] = Prop{Name: "cardnum", Value: cardnum}
	paymentProp[1] = Prop{Name: "country", Value: country}
	p.Properties = paymentProp
}

// prop для платежа на карту приват банка
// bCardOrAcc Карта или счёт получателя
// amt Сумма Напр.: 23.05
// ccy Валюта (UAH, EUR, USD)
// details Назначение платежа
func (p *Payment) SetPrivatPaymentProperties(bCardOrAcc string, amt string, ccy string, details string) {
	paymentProp := make([]Prop, 4)
	paymentProp[0] = Prop{Name: "b_card_or_acc", Value: bCardOrAcc}
	paymentProp[1] = Prop{Name: "amt", Value: amt}
	paymentProp[2] = Prop{Name: "ccy", Value: ccy}
	paymentProp[3] = Prop{Name: "amt", Value: details}
	p.Properties = paymentProp
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
