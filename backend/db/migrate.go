package main

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

func init() {

	utils.InitializeLogger();
}

func main(){
	config.ConnectToDB();
	err := config.DB.AutoMigrate(&models.Note{},&models.User{}).Error;

	if err != nil{
		utils.SugarLogger.Error(err);
	}
}