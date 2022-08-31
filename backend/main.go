package main

import (
	Routers "backend/routers"
	"backend/utils"

	"go.uber.org/zap"
)
func main(){
	utils.InitializeLogger()
	utils.SugarLogger.Info("Hlllllweeer")
	utils.SugarLogger.Error("Not able to log.",zap.String("url", "codewithmukesh.com"))
	r := Routers.SetupRouter()
	
	r.Run()
	
}