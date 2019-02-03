package privat

import "encoding/xml"

type OrdersRequestXML struct {
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

type OrdersResponseXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data struct {
		XMLName xml.Name `xml:"data"`
		Oper    string   `xml:"oper"`
		Info struct {
			XMLName xml.Name `xml:"info"`
			Statements struct {
				XMLName   xml.Name    `xml:"statements"`
				Status    string      `xml:"status,attr"`
				Credit    string      `xml:"credit,attr"`
				Debet     string      `xml:"debet,attr"`
				Statement []Statement `xml:"statement"`
			}
		}
	}
}

type Statement struct {
	XMLName     xml.Name `xml:"statement"`
	Card        string   `xml:"card,attr"`
	Appcode     string   `xml:"appcode,attr"`
	Trandate    string   `xml:"trandate,attr"`
	Amount      string   `xml:"amount,attr"`
	Cardamount  string   `xml:"cardamount,attr"`
	Rest        string   `xml:"rest,attr"`
	Terminal    string   `xml:"terminal,attr"`
	Description string   `xml:"description,attr"`
}

//Выписки по счёту мерчанта - физлица