package common

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	client *http.Client
	once   sync.Once
)

func InitHttpClient(clientConf HttpClient) {
	once.Do(func() {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: func(*http.Request) (*url.URL, error) {
					if clientConf.Proxy == "" {
						return nil, nil
					} else {
						return url.Parse(clientConf.Proxy)
					}
				},
			},
			Timeout: time.Duration(clientConf.Timeout) * time.Millisecond,
		}
	})
}

func HttpGet(url string) (resp *http.Response, err error) {
	if client == nil {
		InitHttpClient(HttpClient{Timeout: 3_000})
	}
	return client.Get(url)
}
