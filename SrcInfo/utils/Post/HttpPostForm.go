package Post

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"test/utils/Headers"
	"time"
)
func HttpPostForm(requestUrl string, header string) (string,string){
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
	client := &http.Client{Transport: Tr}
	data := make(url.Values)
	data.Add("s", "3")
	data.Add("p", "")
	data.Add("sort", "1")
	data.Add("token", "")
	formValue := data.Encode()
	request,_ := http.NewRequest(http.MethodPost, requestUrl, strings.NewReader(formValue))

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
