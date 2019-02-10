package privat

import (
	"encoding/xml"
	"log"
)

type GoldResponse struct {
	Exchangerate struct {
		XMLName      xml.Name `xml:"exchangerate"`
		Exchangerate []GoldExchangerate `xml:"exchangerate"`
	}
}

type GoldExchangerate struct {
	XMLName   xml.Name `xml:"exchangerate"`
	Ccy       string   `xml:"ccy,attr"`         //Код валюты (справочник кодов валют: https://ru.wikipedia.org/wiki/Коды_валют)
	CcyNameRu string   `xml:"ccy_name_ru,attr"` //Название валюты на русском языке
	CcyNameUa string   `xml:"ccy_name_ua,attr"` //Название валюты на украинском языке
	CcyNameEn string   `xml:"ccy_name_en,attr"` //Название валюты на английском языке
	BaseCcy   string   `xml:"base_ccy,attr"`    //Код выбранной вами страны
	Buy       string   `xml:"buy,attr"`         //Курс покупки (коп. * 100)
	Unit      string   `xml:"unit,attr"`        //Количество единиц валюты, которые можно купить по курсу покупки
	Date      string   `xml:"date,attr"`        //Дата последнего обновления курсов валют
}

// Курсы валют, драгоценных металлов НБУ и ЦБ РФ
// country возможные значения :
// ua -курсы НБУ по основным валютам (RUB, USD,EUR)
// ua&full - курсы НБУ по всем валютам, включая драгоценные металлы
// ru - курсы ЦБРФ по основным валютам (USD, EUR)
// ru&full - курсы ЦБ РФ по всем валютам, включая драгоценные металлы
func (api *Privat24Api) GetGold(country string) GoldResponse {
	url := "https://privat24.privatbank.ua/p24api/accountorder?oper=prp&PUREXML&apicour&country="+country

	response, err := api.requestXML(url, nil, "GET")
	if err != nil {
		log.Println(err.Error())
	}

	var goldResponse GoldResponse
	err = xml.Unmarshal(response, &goldResponse)
	if err != nil {
		log.Println(err.Error())
	}

	return goldResponse
}
