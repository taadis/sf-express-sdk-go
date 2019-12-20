package sdk

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
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

	// verifyCode
	xmlWithCheckCodeBuffer := bytes.NewBufferString(xmlRequest)
	xmlWithCheckCodeBuffer.WriteString(checkCode)
	verifyCode := computeVerifyCode(xmlWithCheckCodeBuffer.String())

	// postData
	postDataBuffer := bytes.NewBufferString("xml=")
	postDataBuffer.WriteString(xmlRequest)
	postDataBuffer.WriteString("&verifyCode=")
	postDataBuffer.WriteString(verifyCode)

	req, err := http.NewRequest(
		"POST",     //method string,
		client.url, //url string,
		strings.NewReader(postDataBuffer.String()), //body io.Reader,
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

// 生成认证码
func computeVerifyCode(s string) (verifyCode string) {
	buffer := bytes.NewBufferString(s)
	md5Bytes := md5.Sum(buffer.Bytes())
	verifyCode = base64.StdEncoding.EncodeToString(md5Bytes[:])
	return verifyCode
}
