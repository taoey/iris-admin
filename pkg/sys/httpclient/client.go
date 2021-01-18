/**
 * @Author: taoey
 * @Description:
 * @File:  client
 * @Date: 2020/12/1 11:42
 */
package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func HTTPClient(dailTimeout time.Duration) *http.Client {
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	return &http.Client{
		Transport: &transport,
		Timeout:   dailTimeout,
	}
}

func httpGet(url string, p map[string]string) (string, error) {
	url += "?"
	for key, value := range p {
		url = url + key + "=" + value + "&"
	}
	//do not use default HTTP client!
	var netClient = HTTPClient(time.Second * 10)
	resp, err := netClient.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return "", fmt.Errorf("HttpGet error: %v", err)
	}
	if 200 != resp.StatusCode {
		return "", fmt.Errorf("HttpGet error: %v", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("HttpGet error: %v", err)
	}

	return string(body), nil
}

func HttpPost(url string, contentType string, postbody string) (string, error) {
	return HttpPostWithTimeout(url, contentType, postbody, time.Second*10)
}

func HttpPostWithTimeout(url string, contentType string, postbody string, timeout time.Duration) (string, error) {
	//do not use default HTTP client!
	var netClient = HTTPClient(timeout)

	resp, err := netClient.Post(url, contentType, strings.NewReader(postbody))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", fmt.Errorf("httpPost error: %v", err)
	}
	if 200 != resp.StatusCode {
		return "", fmt.Errorf("httpPost error: %v", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("httpPost error: %v", err)
	}
	return string(body), nil
}

func HttpDownload(url string, filePath string) error {
	regx := "^((https|http|ftp)?://)"
	match, _ := regexp.MatchString(regx, url)
	if !match {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("HttpDownload error: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HttpDownload error: %v", err)
	}

	err = ioutil.WriteFile(filePath, data, 0666) //写入文件
	if err != nil {
		return fmt.Errorf("HttpDownload error: %v", err)
	}
	return nil
}

func HttpDownloadStream(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("HttpDownload error: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HttpDownload error: %v", err)
	}
	return data, nil
}

func Upload(filepath string, extraParams map[string]string, uri string, fileField string) (string, error) {
	request, err := newfileUploadRequest(uri, extraParams, fileField, filepath)
	if err != nil {
		return "", err
	}
	client := HTTPClient(30 * time.Second)
	resp, err := client.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return body.String(), nil
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
