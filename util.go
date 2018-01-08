package goqiwi

import (
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

func addOptions(baseUrl string, opt interface{}) (string, error) {
	value := reflect.ValueOf(opt)
	if value.Kind() == reflect.Ptr && value.IsNil() || value.Kind() == reflect.Invalid {
		return baseUrl, nil
	}

	URL, err := url.Parse(baseUrl)
	if err != nil {
		return baseUrl, err
	}

	queryString, err := query.Values(opt)
	if err != nil {
		return baseUrl, err
	}

	URL.RawQuery = queryString.Encode()
	return URL.String(), nil
}
