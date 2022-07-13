package Get

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"test/utils/Headers"
	"time"
)

func HttpGet(requestUrl string, header string) (string,string){
	defer func(){
		err := recover()
		_ = err
	}()
	key,val := Headers.AutoSetHeaders(header)
	Tr := &http.Transport{
		MaxIdleConns: 5,
		ExpectContinueTimeout: 2000 * time.Millisecond,
		DialContext: (&net.Dialer{
			Timeout: 2000 * time.Millisecond,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: Tr,Timeout: 5*time.Second}
	request,_ := http.NewRequest(http.MethodGet, requestUrl, nil)
	for count,i := range key{
		request.Header.Add(i, val[count])
	}

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
