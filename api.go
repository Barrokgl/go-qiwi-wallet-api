package goqiwi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const apiLink = "https://edge.qiwi.com/"

type QiwiApi struct {
	Token  string
	Client *http.Client
}

func NewQiwiApi(token string, client *http.Client) *QiwiApi {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 30,
		}
	}

	return &QiwiApi{
		Token:  token,
		Client: client,
	}
}

// get qiwi profile information
func (api *QiwiApi) GetProfile(params ProfileParams) (*Profile, error) {
	baseUrl := apiLink + "person-profile/v1/profile/current"

	URL, err := addOptions(baseUrl, params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result Profile
	json.Unmarshal(body, &result)
	return &result, nil
}

// get full history of payments
func (api *QiwiApi) GetHistory(wallet string, params HistoryParams) (*History, error) {
	if string([]rune(wallet)[0]) == "+" {
		wallet = wallet[:len(wallet) - len("+")]
	}

	baseUrl := apiLink + "/payment-history/v1/persons/" + wallet + "/payments"

	URL, err := addOptions(baseUrl, params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result History
	json.Unmarshal(body, &result)
	return &result, nil
}

// get statistic of payments by period
func (api *QiwiApi) GetPaymentStatistic(wallet string, params PaymentStatisticParams) (*PaymentStatistic, error) {
	if string([]rune(wallet)[0]) == "+" {
		wallet = wallet[:len(wallet) - len("+")]
	}

	baseUrl := apiLink + "/payment-history/v1/persons/" + wallet + "/total"

	URL, err := addOptions(baseUrl, params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result PaymentStatistic
	json.Unmarshal(body, &result)
	return &result, nil
}

// get balance of your account
func (api *QiwiApi) GetBalance() (*Balances, error) {
	baseUrl := apiLink + "funding-sources/v1/accounts/current"

	body, err := api.request(baseUrl, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result Balances
	json.Unmarshal(body, &result)
	return &result, nil
}

// get standard rate by provider code
func (api *QiwiApi) GetStandardRate(providerCode string) (*StandardRate, error) {
	baseUrl := apiLink + "sinap/providers/" + providerCode + "/form"

	body, err := api.request(baseUrl, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result StandardRate
	json.Unmarshal(body, &result)
	return &result, nil
}

// get special rate by provider code
func (api *QiwiApi) GetSpecialRate(providerCode string, params SpecialRateParams) (*SpecialRate, error) {
	baseUrl := apiLink + "sinap/providers/" + providerCode + "/onlineCommission"

	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(baseUrl, "POST", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var result SpecialRate
	json.Unmarshal(body, &result)
	return &result, nil
}

// basic request
func (api *QiwiApi) request(URL, method string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+api.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return result, err
}
