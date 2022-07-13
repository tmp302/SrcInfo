package Get

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"test/utils/DB"
	"test/utils/Decoding"
	"test/utils/Message"
)
func GetVbNew(){
	var endStatus = 0
	var newProject = ""
	var newLen []int
	var vulBoxUrl = "https://vbapi.vulbox.com/v1/task/getList?search_value=&page=1&order=default&limit=100&sort=ascend&task_type=2&trade_type=&skill_type=&business_type=&other_type="
	questInfo,errInfo := HttpGet(vulBoxUrl,
		`GET / HTTP/1.1
Host: vbapi.vulbox.com
Sec-Ch-Ua: " Not;A Brand";v="99", "Microsoft Edge";v="103", "Chromium";v="103"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "Windows"
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6
Connection: close`)
	if errInfo == "" {
		test := strings.ReplaceAll(Decoding.DeUnicode(questInfo), `\/`, `/`)
		resultA := gjson.Get(test, "data").Get("list")
		//fmt.Println(resultA)

		for i := range resultA.Array() {
			id := gjson.Get(resultA.Array()[i].String(), "id").Int()
			taskTypeName := gjson.Get(resultA.Array()[i].String(), "task_type_name").String()
			taskTitle := gjson.Get(resultA.Array()[i].String(), "task_title").String()
			taskStime := gjson.Get(resultA.Array()[i].String(), "task_stime").String()
			taskEtime := gjson.Get(resultA.Array()[i].String(), "task_etime").String()
			if DB.SearchExist(int(id), "VulBox") == 0 {
				DB.InsertVulBox(int(id), taskTypeName, taskTitle, taskStime, taskEtime)
				newProject += strconv.Itoa(int(id)) + "\t" + taskTypeName + "\t" + taskTitle + "\t" + taskStime + "\t" + taskEtime + "\n"
				endStatus = 1
				newLen = append(newLen, len(newProject))
			}
		}
		if newProject != "" {
			if len(newLen) > 5{
				fmt.Println("[!] 新项目数量过多，不进行推送 --VulBox")
				fmt.Println("<==================================================")
				fmt.Println(newProject + "==================================================>")
			} else
			{
				Message.SendMessage("盒子", newProject)
			}
		}
		if endStatus == 0 {
			fmt.Println("暂无新厂商 --VulBox")
		}
	}else{
		fmt.Println(errInfo)
	}

}
