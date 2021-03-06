package main

import (
	"example/web-service-gin/Config"
	"example/web-service-gin/Models"
	"example/web-service-gin/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Person{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
