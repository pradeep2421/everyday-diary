package controllers

import (
	"backend/controllers/services"
	"backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)



func CreateUser(c *gin.Context) {
	var user models.User;
	c.Bind(&user);
	err := services.CreateUser(&user);

	if(err != nil){
		c.AbortWithStatusJSON(400,gin.H{"message": "cannot create user","error":err});
		return;
	}
	c.JSON(200,gin.H{"message":"user created successfull","user":user});
	
}
func LoginUser(c *gin.Context) {
	var login models.Login;
	
	c.Bind(&login);
	user ,err := services.LoginUser(&login);
	if err != nil{
		c.AbortWithStatusJSON(400,gin.H{"message":"database error","error":err});
		return;
	}

	validUser := (user.Name != ""); 
	if validUser {
		c.JSON(200,gin.H{"success":validUser, "message":"you are logged in!!","user":user});
		return;
	}
	c.JSON(200, gin.H{"success":validUser,"message":"name or password  is incorrect"});


}

func ViewUser(c *gin.Context) {
	var users []models.User;
	err = services.ViewUser(&users);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":err});
		return;
	}
	c.JSON(200, gin.H{"users":users});
	
	
}
func UpdateUser(c *gin.Context) {
	var user models.User;
	id := c.Param("id");
	newId,_ := strconv.ParseUint(id,10,64);
	user.ID = uint(newId);

	c.Bind(&user);

	err = services.UpdateUser(&user);
	if err != nil{
		c.AbortWithStatusJSON(400,gin.H{"message":"database error","error":err});
		return;
	}
	c.JSON(200, gin.H{"message":"user updated","user":user});
}

func SingleUser(c *gin.Context) {
	var user models.User;
	id := c.Param("id");
	err  = services.SingleUser(&user,id);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":err});
		return;
	}
	c.JSON(200, gin.H{"user":user});
	
}

func DeleteUser(c *gin.Context) {
	//get the id off the url
	id := c.Param("id");
	
	err := services.DeleteUser(id);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":"cannot delete user","error":err});
		return;
	}
	c.JSON(200,gin.H{"message":"user deleted"});
}