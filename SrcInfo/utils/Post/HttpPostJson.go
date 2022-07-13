package Post

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)
func HttpPostJson(requestUrl string, jsonParams interface{}) (string,string){
	defer func(){
		err := recover()
		_ = err
	}()

	//proxyUrl,_ := url.Parse("http://127.0.0.1:8080")
	Tr := &http.Transport{
		//Proxy: http.ProxyURL(proxyUrl),
		DialContext: (&net.Dialer{
			Timeout: 2000 * time.Millisecond,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: Tr}
	payload,_ := json.Marshal(jsonParams)
	request,_ := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(payload))
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36")

	startRequest,requestDoError := client.Do(request)
	if requestDoError != nil {
		return "", "请求错误: " + requestDoError.Error()
	}
	if startRequest.StatusCode != 200{
		return "", "响应状态错误: " + strconv.Itoa(startRequest.StatusCode)
	}

	defer func(){_= startRequest.Body.Close()}()
	bodyText, ioReadError := ioutil.ReadAll(startRequest.Body)
	if ioReadError != nil{
		return "", "读取错误: " + ioReadError.Error()
	}

	return string(bodyText), ""
}
