package goqiwi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const apiLink = "https://edge.qiwi.com/"

type QiwiApi struct {
	Token  string
	Client *http.Client
}

func NewQiwiApi(token string, client *http.Client) *QiwiApi {
	if client == nil {
		client = &http.Client{}
	}

	return &QiwiApi{
		Token:  token,
		Client: client,
	}
}

// get qiwi profile information
func (api *QiwiApi) GetProfile(params GetProfileParams) (*GetProfileResult, error) {
	baseUrl := apiLink + "person-profile/v1/profile/current"

	URL, err := addOptions(baseUrl, params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result GetProfileResult
	json.Unmarshal(body, &result)
	return &result, nil
}

// get full history of payments
func (api *QiwiApi) GetHistory(wallet string, params GetHistoryParams) (*GetHistoryResult, error) {
	baseUrl := apiLink + "/payment-history/v1/persons/" + wallet + "/payments"

	URL, err := addOptions(baseUrl, params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result GetHistoryResult
	json.Unmarshal(body, &result)
	return &result, nil
}

// get statistic of payments by period
func (api *QiwiApi) GetPaymentStatistic(wallet string, params GetPaymentStatisticParams) (*GetPaymentStatisticResult, error) {
	baseUrl := apiLink + "/payment-history/v1/persons/" + wallet + "/total"

	URL, err := addOptions(baseUrl, params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", nil)
	if err != nil {
		return nil, err
	}

	var result GetPaymentStatisticResult
	json.Unmarshal(body, &result)
	return &result, nil
}

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
