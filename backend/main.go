package main

import (
	"backend/config"
	"backend/models"
	Routers "backend/routers"
	"backend/utils"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)
func init(){ 
	utils.InitializeLogger()

}
var err error
func main(){
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		utils.SugarLogger.Error("DataBase is not connected",zap.Error(err))
	}
 err = config.DB.AutoMigrate(&models.User{}).Error
 if err != nil {
	utils.SugarLogger.Error("DataBase is not connected",zap.Error(err))
	return
}
	r := Routers.SetupRouter()
	
	r.Run()
	
}