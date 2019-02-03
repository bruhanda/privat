package privat

import (
	"net/http"
	"log"
	"time"
	"errors"
	"io"
	"io/ioutil"
	"github.com/golang/mock/mockgen/tests/custom_package_name/client/v1"
	"net/rpc"
)

const apiUrl = "https://api.privatbank.ua"

type Privat24Api struct {
	merchantID       int
	merchantPassword string
	client           *http.Client
	apiUrl           string
}

func NewApi(merchantID int, merchantPassword string) (*Privat24Api, error) {
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	api := &Privat24Api{merchantID: merchantID, merchantPassword:merchantPassword, client: client, apiUrl: apiUrl}

	return api, nil
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
