package goqiwi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/net/proxy"
)

// default qiwi api link
const apiLink = "https://edge.qiwi.com/"

type QiwiApi struct {
	token  string
	client *http.Client
	apiUrl string
	uuid   string
}

// returns QiwiApi instance
// proxyAdd - socks5 proxy address, format: 127.0.0.1:1234
func NewQiwiApi(token, apiUrl, proxyAddr string, auth proxy.Auth) (*QiwiApi, error) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	if apiUrl == "" {
		apiUrl = apiLink
	}

	api := &QiwiApi{token: token, client: client, apiUrl: apiUrl}

	if proxyAddr != "" {
		api.setHTTPProxy(proxyAddr, auth)
	}

	api.uuid = uuid.NewV4().String()

	return api, nil
}

// sets new socks5 proxy
func (api *QiwiApi) SetSOCKS5(proxyAddr string, auth proxy.Auth) error {
	dialer, err := proxy.SOCKS5("tcp", proxyAddr, &auth, proxy.Direct)
	if err != nil {
		return err
	}
	httpTransport := &http.Transport{}
	api.client.Transport = httpTransport
	httpTransport.Dial = dialer.Dial

	return nil
}

func (api *QiwiApi) setHTTPProxy(proxyAddr string, auth proxy.Auth) error {
	proxyURL, err := url.Parse(proxyAddr)
	if err != nil {
		return err
	}
	headers := make(http.Header)
	headers.Add("Proxy-Key", api.uuid)

	api.client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL), TLSClientConfig: &tls.Config{},
		ProxyConnectHeader: headers,
	}
	return nil
}

// sets new token
func (api *QiwiApi) SetToken(token string) {
	api.token = token
}

// get qiwi profile information
func (api *QiwiApi) GetProfile(params ProfileParams) (*Profile, error) {
	baseUrl := api.apiUrl + "person-profile/v1/profile/current"

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
		wallet = wallet[:len(wallet)-len("+")]
	}

	baseUrl := api.apiUrl + "payment-history/v1/persons/" + wallet + "/payments"

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
		wallet = wallet[:len(wallet)-len("+")]
	}

	baseUrl := api.apiUrl + "/payment-history/v1/persons/" + wallet + "/total"

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
	baseUrl := api.apiUrl + "funding-sources/v1/accounts/current"

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
	baseUrl := api.apiUrl + "sinap/providers/" + providerCode + "/form"

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
	baseUrl := api.apiUrl + "sinap/providers/" + providerCode + "/onlineCommission"

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

// make payment from your account
func (api *QiwiApi) Payment(providerCode, account string, amount float64) (*Payment, error) {
	params := PaymentParams{
		ID:            strconv.FormatInt(time.Now().UnixNano(), 10),
		Sum:           Sum{Amount: amount, Currency: "643"},
		Source:        "account_643",
		PaymentMethod: PaymentMethod{Type: "Account", AccountId: "643"},
		Fields:        Fields{Account: account},
	}

	baseUrl := api.apiUrl + "sinap/api/v2/terms/" + providerCode + "/payments"

	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(baseUrl, "POST", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var result Payment
	json.Unmarshal(body, &result)
	return &result, nil
}

// make payment on Qiwi
// same as (api *QiwiApi) Payment , but with hardcoded provider code
func (api *QiwiApi) PaymentQiwi(phone, comment string, amount float64) (*Payment, error) {
	params := PaymentParams{
		ID:            strconv.FormatInt(time.Now().UnixNano(), 10),
		Sum:           Sum{Amount: amount, Currency: "643"},
		Source:        "account_643",
		PaymentMethod: PaymentMethod{Type: "Account", AccountId: "643"},
		Comment:       comment,
		Fields:        Fields{Account: phone},
	}
	baseUrl := api.apiUrl + "sinap/api/v2/terms/99/payments"

	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err := api.request(baseUrl, "POST", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var result Payment
	json.Unmarshal(body, &result)
	return &result, nil
}

// determines mobile operator code
func (api *QiwiApi) DetermineOperator(phone string) (*DeterminedProvider, error) {
	baseUrl := "https://qiwi.com/mobile/detect.action?phone=" + phone

	body, err := api.request(baseUrl, "POST", nil)
	if err != nil {
		return nil, err
	}

	var result DeterminedProvider
	json.Unmarshal(body, &result)
	return &result, nil
}

// determines card provider
func (api *QiwiApi) DetermineCard(cardNumber string) (*DeterminedProvider, error) {
	baseUrl := "https://qiwi.com/card/detect.action?cardNumber=" + cardNumber

	body, err := api.request(baseUrl, "POST", nil)
	if err != nil {
		return nil, err
	}

	var result DeterminedProvider
	json.Unmarshal(body, &result)
	return &result, nil
}

// basic request
func (api *QiwiApi) request(URL, method string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+api.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

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
