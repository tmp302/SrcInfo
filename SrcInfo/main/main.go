package main

import (
	"test/utils/DB"
	"test/utils/Get"
)

func main(){
	DB.CreateDB()
	Get.GetVbNew()
	Get.GetBtNew()
}
