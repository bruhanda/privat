package privat

import "encoding/xml"

type VisaPaymentRequestXML struct {
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

type VisaPaymentResponseXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data struct {
		XMLName xml.Name `xml:"data"`
		Oper    string   `xml:"oper"`
		Payment struct {
			XMLName  xml.Name `xml:"payment"`
			ID       string   `xml:"id,attr"`      //Уникальный идентификатор платежа, присвоенный партнером платежей
			State    string   `xml:"state,attr"`   //Состояние платежа (1 - проведён, 0- забракован)
			Message  string   `xml:"message,attr"` //Расширенное сообщение о состоянии платежа
			Ref      string   `xml:"ref,attr"`     //Идентификатор платежа в Приват24
			Amt      string   `xml:"amt,attr"`     //Сумма платежа (без комиссии)
			Ccy      string   `xml:"ccy,attr"`     //Валюта операции
			Comis    string   `xml:"comis,attr"`   //Сумма комиссии банка по данному типу платежа
			Code     string   `xml:"code,attr"`    //Код ваучера (для prepaid-операций)
			Cardinfo string   `xml:"cardinfo,attr"`
		}
	}
}