package privat

import (
	"net/http"
	"log"
	"time"
	"errors"
	"io"
	"io/ioutil"
	"encoding/xml"
)

const apiUrl = "https://api.privatbank.ua/p24api/"

type Privat24Api struct {
	merchantID       int
	merchantPassword string
	client           *http.Client
	apiUrl           string
}

func NewPublicApi() *Privat24Api {
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	api := &Privat24Api{client: client, apiUrl: apiUrl}

	return api
}

func NewApi(merchantID int, merchantPassword string) *Privat24Api {
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	api := &Privat24Api{merchantID: merchantID, merchantPassword: merchantPassword, client: client, apiUrl: apiUrl}

	return api
}

func (api *Privat24Api) requestXML(url string, body io.Reader, method string) ([]byte, error) {
	req, err := http.NewRequest(http.Request{}.Method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	response, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		log.Println("[REQUEST ERROR]: ", response.Status, string(result))
		err = errors.New(response.Status + " : " + string(result))
	}

	return result, err
}

func (api *Privat24Api) getMerchantStruct(data interface{}) Merchant {
	res, err := xml.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	merchant := new(Merchant)
	merchant.ID = api.merchantID
	merchant.Signature = SHA1(GetMD5Hash(string(res) + api.merchantPassword))

	return *merchant
}
