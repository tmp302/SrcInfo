package DB

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func SearchExist(id int, dbName string) int{
	db,err := sql.Open("sqlite3", "./SrcInfo.db")
	if err != nil{panic(err)}

	var tmpTest int
	db.QueryRow("select Id from "+dbName+" where Id="+strconv.Itoa(id)).Scan(&tmpTest)
	if id == tmpTest{
		return 1
	}else{
		return 0
	}
}
