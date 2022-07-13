package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func InsertVulBox(Id int, ProjectType string, Name string, StartTime string, EndTime string){
	resultTxt := strconv.Itoa(Id) + " " + Name
	if SearchExist(Id, "VulBox") == 0 {
		db,_ := sql.Open("sqlite3", "./SrcInfo.db")
		inVulBoxSql, _ := db.Prepare("insert into VulBox values(?,?,?,?,?)")
		_, err := inVulBoxSql.Exec(Id, ProjectType, Name, StartTime, EndTime)
		if err != nil {panic(err)}
		fmt.Println("[+] " + resultTxt + "写入成功 --VulBox")
	}else{
		fmt.Println("[-] " + resultTxt + "已存在，不进行写入 --VulBox")
	}
}
func InsertBuTian(Id int, Name string, MaxReward int){
	resultTxt := strconv.Itoa(Id) + " " + Name
	if SearchExist(Id, "BuTian") == 0 {
		db, _ := sql.Open("sqlite3", "./SrcInfo.db")
		inBuTianSql, _ := db.Prepare("insert into BuTian values(?,?,?)")
		_, err := inBuTianSql.Exec(Id, Name, MaxReward)
		if err != nil {panic(err)}

		fmt.Println("[+] " + resultTxt + "写入成功 --BuTian")
	}else{
		fmt.Println("[-] " + resultTxt + "已经存在，不进行写入 --BuTian")
	}
}