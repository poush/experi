/*
* @Author: Piyush Agrawal
* @Date:   2017-10-23 20:47:54
* @Last Modified by:   Piyush Agrawal
* @Last Modified time: 2017-10-23 22:46:55
*/
package controllers

import "github.com/revel/revel"
import "github.com/revel/modules/gorm/app"
import "os"
import "net"
import "net/url"
import "github.com/poush/sharkDeploy/app/models"

func InitializeDB() {
    s := os.Getenv("DATABASE_URL")
    u, _ := url.Parse(s)
	params := gorm.DbInfo{}

	host, _, _ := net.SplitHostPort(u.Host)
	params.DbDriver = "postgres"
	params.DbHost = host
	params.DbUser = u.User.Username()
	params.DbPassword, _ = u.User.Password()
	params.DbPassword = params.DbPassword + "  sslmode=require"
	params.DbName = u.Path[1:]
	gorm.InitDBWithParameters(params)
	gorm.DB.AutoMigrate(&models.User{})
}


func init() {
	revel.OnAppStart(func() {
		InitializeDB()
	})
}