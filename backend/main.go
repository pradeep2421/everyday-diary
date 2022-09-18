package main

import (
	"backend/config"
	"backend/initializers"
	"backend/routers"
	"backend/utils"

	_ "github.com/go-sql-driver/mysql"
)
func init(){ 
	utils.InitializeLogger()
	initializers.LoadEnvVariables();

}


func main(){
	config.ConnectToDB(); 
	r := routers.SetupRouter()

	r.Run()
	
}