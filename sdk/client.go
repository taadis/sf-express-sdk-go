package sdk

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// 客户端结构体
type Client struct {
	httpClient *http.Client
	url        string
}

// 实例化一个新客户端
func NewClient() (client *Client, err error) {
	client = &Client{}
	client.httpClient = &http.Client{}
	client.url = "https://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService"
	return client, nil
}

// 干活
func (client *Client) ExpressService(xmlRequest string, checkCode string) (xmlResponse string, err error) {
	req, err := http.NewRequest(
		"POST",                        //method string,
		client.url,                    //url string,
		strings.NewReader(xmlRequest), //body io.Reader,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	res, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	xmlResponse = string(body[:])
	return xmlResponse, nil
}
