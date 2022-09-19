package middleware

import (
	"backend/config"
	"backend/controllers"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context){
	
	//get the cookie 
	cookie ,err := c.Cookie("jwt");
	if err != nil {
		c.AbortWithStatusJSON(400,gin.H{"message":"token not found"});
		return;
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(controllers.SecretKey), nil
	})
	// fmt.Println(token,err);
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64){
		
			c.AbortWithStatus(http.StatusUnauthorized);
		}
		var user models.User
		config.DB.First(&user,claims["Issuer"]);
	
		id ,_:= strconv.ParseUint( c.Param("id"),10,64);
		
		if( user.ID != uint(id) ){
		
			c.AbortWithStatus(http.StatusUnauthorized);

		}
		c.Set("user",user);
	
		c.Next();
	} else {
		fmt.Println("NO Authorization");
		fmt.Println(err)
	}




}