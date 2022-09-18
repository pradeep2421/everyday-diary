package controllers

import (
	"backend/controllers/services"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)



func RegisterUser(c *gin.Context) {
	
	
}
func LoginUser(c *gin.Context) {
	
	fmt.Println("Hello From Login")
	// fmt.Println("Hello From Login")
	services.CreateUser();
	utils.SugarLogger.Warn("This is a Drile")
}

func GetUserProfile(c *gin.Context) {
	
	
}
func EditUserProfile(c *gin.Context) {
	
	
}

