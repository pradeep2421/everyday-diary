package main

import (
	"backend/config"
	"backend/initializers"
	"backend/models"
	"backend/utils"
)

func init() {
	initializers.LoadEnvVariables();

	utils.InitializeLogger();
}

func main(){
	config.ConnectToDB();
	err := config.DB.AutoMigrate(&models.Note{},&models.User{}).Error;

	if err != nil{
		utils.SugarLogger.Error(err);
	}
}