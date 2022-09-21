package main

import (
	"backend/config"
	"backend/routers"
	"backend/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
func init(){ 
	utils.InitializeLogger()

}


func main(){
	viper_config, err := utils.LoadConfig("./utils");

	if err !=nil {
		utils.SugarLogger.Error("error in loading config file with viper");
	}
	fmt.Println(viper_config.Db.DbDriver);
	config.ConnectToDB(); 
	r := routers.SetupRouter()

	r.Run(":3000")
	
}