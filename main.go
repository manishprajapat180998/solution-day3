package main

import (
	"fmt"
	"github.com/exercise2/Config"
	"github.com/exercise2/Models"
	"github.com/exercise2/Routes"
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