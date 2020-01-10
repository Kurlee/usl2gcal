package webClient

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"log"
	"net/http"
	"time"
)

func Get_https(url string) *otto.Value {
	tr := &http.Transport{
		Proxy:                  nil,
		DialContext:            nil,
		Dial:                   nil,
		DialTLS:                nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     true, // changed
		MaxIdleConns:           10, // changed
		MaxIdleConnsPerHost:    0,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        30 * time.Second, // changed
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}

	client := &http.Client{
		Transport:     tr, // changed
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	resp, err := client.Get(url)
	check(err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	firstScript := doc.Find("script").First()

	vm := otto.New()

	_, err = vm.Run("var window = {};")
	check(err)

	_, err = vm.Run(firstScript.Text())
	check(err)

	wVal, err := vm.Get("window")
	check(err)

	pdata, err := getValueFromObject(wVal, "POST_DATA")
	check(err)

	return pdata
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getValueFromObject(val otto.Value, key string) (*otto.Value, error) {
	if !val.IsObject() {
		return nil, errors.New("passed val is not an Object")
	}

	valObj := val.Object()

	obj, err := valObj.Get(key)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

