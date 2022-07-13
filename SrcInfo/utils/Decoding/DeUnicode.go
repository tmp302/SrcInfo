package Decoding

import (
	"strconv"
	"strings"
)

func DeUnicode(contextStr string) string{
	textStr,_ := strconv.Unquote(strings.ReplaceAll(strconv.Quote(contextStr), `\\u`, `\u`))
	return textStr
}
