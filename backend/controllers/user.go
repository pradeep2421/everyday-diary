package controllers

import (
	"backend/controllers/services"
	"backend/controllers/validation"
	"backend/models"
	"backend/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = "dracarys";

func CreateUser(c *gin.Context) {
	var user models.User;
	c.Bind(&user);
	password ,_:= bcrypt.GenerateFromPassword([]byte(user.Password),14);

	
	user.Password = string(password)
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

	if err = validation.LoginValidation(login);err != nil{
		fmt.Println(err);
	};

	
	user ,err := services.LoginUser(&login);
	if err != nil{
		c.AbortWithStatusJSON(400,gin.H{"message":"database error","error":err});
		return;
	}
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Issuer": strconv.Itoa(int(user.ID)),
			"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
		})
		token, err1 := claims.SignedString([]byte(SecretKey));
		if (err1 != nil){
			utils.SugarLogger.Error(err1);
			c.AbortWithStatusJSON(400,gin.H{"message":"Internal Server Error","error":err1});
		return;
		}
		c.SetSameSite(http.SameSiteLaxMode);
		c.SetCookie("jwt", token, 60*60*24, "", "", false, true)


		c.JSON(200,gin.H{"success":true, "message":"you are logged in!!","user":user});

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