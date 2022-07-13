package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"test/utils/Read"
)

func CreateDB(){
	if Read.ReadConfig("DBCreate.status") == "0" {
		var vulBoxInfoCS int
		var buTianInfoCS int

		db, err := sql.Open("sqlite3", "./SrcInfo.db")
		if err != nil {
			panic(err)
		}

		vulBoxInfoVerify := db.QueryRow("select count(name) from sqlite_master where name='VulBox';").Scan(&vulBoxInfoCS)
		if vulBoxInfoVerify != nil {
			log.Fatal(vulBoxInfoVerify)
		}
		buTianInfoVerify := db.QueryRow("select count(name) from sqlite_master where name='BuTian';").Scan(&buTianInfoCS)
		if buTianInfoVerify != nil {
			log.Fatal(buTianInfoVerify)
		}

		switch vulBoxInfoCS {
		case 0:
			createVulBoxDB := `create table VulBox(
			Id int PRIMARY KEY,
			ProjectType varchar(50),
			Name varchar(50),
			StartTime varchar(50),
			EndTime varchar(50)
			);`
			db.Exec(createVulBoxDB)
			fmt.Println("VulBox 表已创建成功")
		case 1:
			fmt.Println("VulBox 表已存在，不进行创建")
		default:
			fmt.Println("状态异常: ", vulBoxInfoCS)
		}

		switch buTianInfoCS {
		case 0:
			createBuTianDB := `create table BuTian(
			Id int PRIMARY KEY,
			Name varchar(50),
			MaxReward int
			);`
			db.Exec(createBuTianDB)
			fmt.Println("BuTian 表已创建成功")
		case 1:
			fmt.Println("BuTian 表已存在，不进行创建")
		default:
			fmt.Println("状态异常: ", buTianInfoCS)
		}
	}else{
		fmt.Println("SrcInfo.db 文件已存在不进行创建")
	}
}
