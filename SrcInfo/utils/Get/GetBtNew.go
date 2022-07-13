package Get

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"test/utils/DB"
	"test/utils/Decoding"
	"test/utils/Message"
	"test/utils/Post"
)

func GetBtNew(){
	var endStatus = 0
	var newProject = ""
	var newLen []int
	var buTianUrl = "https://www.butian.net/Reward/corps"
	questInfo,errInfo := Post.HttpPostForm(buTianUrl,
		`Post / HTTP/1.1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36
accept-ranges: bytes
access-control-allow-origin: *

`)
	if errInfo == "" {
		test := strings.ReplaceAll(Decoding.DeUnicode(questInfo), `\/`, `/`)
		resultA := gjson.Get(test, "data").Get("list")
		//fmt.Println(resultA)

		for i, _ := range resultA.Array() {
			companyId := gjson.Get(resultA.Array()[i].String(), "company_id").Int()
			companyName := gjson.Get(resultA.Array()[i].String(), "company_name").String()
			maxReward := gjson.Get(resultA.Array()[i].String(), "reward").Int()
			if DB.SearchExist(int(companyId), "BuTian") == 0 {
				DB.InsertBuTian(int(companyId), companyName, int(maxReward))
				newProject += strconv.Itoa(int(companyId)) + "\t" + companyName + "\t" + strconv.Itoa(int(maxReward)) + "\n"
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
			fmt.Println("暂无新厂商 --BuTian")
		}
	}
}
