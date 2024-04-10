package subscription

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

// FetchSubscription 从给定的URL获取订阅内容。
func FetchSubscription(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DecodeBase64 对Base64编码的字符串进行解码。
func DecodeBase64(encoded string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
