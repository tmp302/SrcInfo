package Headers

import "strings"

func AutoSetHeaders(Headers string) ([]string,[]string){
	var (
		keyResult []string
		valResult []string
	)

	requestInfo := Headers
	allInfo:= strings.Split(requestInfo, "\n\n")
	headers := strings.Split(allInfo[0], "\n")

	for _,i:= range headers[1:]{
		key := strings.SplitN(i, ":", 2)
		keyResult = append(keyResult, key[0])
		valResult = append(valResult, key[1])
	}
	return keyResult,valResult
}
