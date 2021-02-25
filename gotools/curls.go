/**********

#### curls 常用网络方法

```
* HttpGet //发送普通Get请求
* HttpParamsGet //发送普通Get带参数请求
* HttpHeadGet //带请求头get请求
* HttpPost //发送post请求
* HttpPostJson //发送JSON数据POST方法一
* HttpPostJson2 //发送JSON数据POST方法二
```

*****/

package gotools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * 发送普通Get请求
 */
func HttpGet(httpUrl string) error {
	res, _ :=http.Get(httpUrl)
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 发送普通Get带参数请求
 */
func HttpParamsGet(httpUrl string, params map[string]string) error {
	urlParse, _:= url.Parse(httpUrl)
	//如果参数中有中文参数,这个方法会进行URLEncode
	urlValue := url.Values{}
	for i, v := range params {
		urlValue.Set(i, v)
	}
	urlParse.RawQuery = urlValue.Encode()
	urlPath := urlParse.String()
	resp,_ := http.Get(urlPath)
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 带请求头get请求
 */
func HttpHeadGet(httpUrl string, heads map[string]string) error {
	client := &http.Client{}
	req,_ := http.NewRequest("GET", httpUrl, nil)
	for i, v := range heads {
		req.Header.Add(i, v)
	}
	resp,_ := client.Do(req)
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 发送post请求
 */
func HttpPost(httpUrl string, params map[string]string) error {
	urlValues := url.Values{}
	for i, v := range params {
		urlValues.Add(i, v)
	}
	resp, _ := http.PostForm(httpUrl, urlValues)
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 发送JSON数据POST方法一
 */
func HttpPostJson(httpUrl string, params map[string]interface{}) error {
	client := &http.Client{}
	bytesData, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", httpUrl, bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 发送JSON数据POST方法二
 */
func HttpPostJson2(httpUrl string, params map[string]interface{}) error {
	bytesData, _ := json.Marshal(params)
	resp, _ := http.Post(httpUrl,"application/json", bytes.NewReader(bytesData))
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}