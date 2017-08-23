package goqiwi

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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

func (api *QiwiApi) GetProfile(params GetProfileParams) (*GetProfileResult, error) {
	URL := apiLink + "person-profile/v1/profile/current"

	value, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(URL, "GET", strings.NewReader(value.Encode()))
	if err != nil {
		return nil, err
	}

	var result GetProfileResult
	json.Unmarshal(body, &result)
	return &result, nil
}

func (api *QiwiApi) GetHistory(wallet string, params GetHistoryParams) error {
	URL := apiLink + "/payment-history/v1/persons/" + wallet + "/payments"

	value, err := query.Values(params)
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
