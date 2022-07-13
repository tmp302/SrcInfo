package Message

import (
	"fmt"
	"test/utils/Post"
	"test/utils/Read"
)

func SendMessage(platform string,content string){

	u := struct{
		Token string `json:"token"`
		Title string `json:"title"`
		Content string `json:"content"`
		Template string `json:"template"`
		Channel string `json:"channel"`
	}{
		Token: Read.ReadConfig("WxMess.api"),
		Title: "新项目" + " --" + platform,
		Content: content,
		Template: "txt",
		Channel: "wechat",
	}
	questInfo,errInfo := Post.HttpPostJson("https://www.pushplus.plus/api/send", u)
	if errInfo == ""{
		fmt.Println(questInfo)
	}else{
		fmt.Println(errInfo)
	}
}