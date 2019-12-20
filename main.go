package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwtapp/Config"
	"jwtapp/Models"
	"jwtapp/Routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()

	// run the migrations:
	Config.DB.AutoMigrate(&Models.User{}).AddForeignKey("authorityId", "authorities(authorityId)", "RESTRICT", "RESTRICT")
	Config.DB.AutoMigrate(&Models.Authority{})
	Config.DB.AutoMigrate(&Models.Product{})
	//Config.DB.AutoMigrate(&Models.ApiAuthority{})




	// Setup routes
	r := Routes.SetupRoutes()
	r.Use(gin.Recovery())
	// running
	_ = r.Run(":9000")

}