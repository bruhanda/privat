package privat

import (
	"encoding/xml"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"log"
)

type QueryXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant
	Data     Data
}

type Merchant struct {
	XMLName   xml.Name `xml:"merchant"`
	ID        int      `xml:"id"`
	Signature string   `xml:"signature"`
}

type Prop struct {
	XMLName xml.Name `xml:"prop"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Payment struct {
	XMLName    xml.Name `xml:"payment"`
	ID         int      `xml:"id,attr"`
	Properties []Prop
}

type Oper struct {
	XMLName xml.Name `xml:"oper"`
}

type Wait struct {
	XMLName xml.Name `xml:"wait"`
}

type Test struct {
	XMLName xml.Name `xml:"test"`
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Items   []interface{}
}

func NewQueryXML(merchantID int, merchantPassword string, data Data) []byte {
	queryXML := new(QueryXML)
	queryXML.Version = "1.0"
	queryXML.Data = data
	queryXML.setMerchant(merchantID, merchantPassword)
	res, err := xml.Marshal(queryXML)
	if err != nil {
		log.Println(err)
	}

	return res
}

func (qd *QueryXML) setMerchant(merchantID int, merchantPassword string) {
	res, err := xml.Marshal(qd.Data)
	if err != nil {
		log.Println(err)
	}
	merchant := new(Merchant)
	merchant.ID = merchantID
	merchant.Signature = SHA1(GetMD5Hash(string(res) + merchantPassword))
	qd.Merchant = *merchant
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
