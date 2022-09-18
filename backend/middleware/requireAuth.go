package middleware

import (
	"backend/config"
	"backend/controllers"
	"backend/models"
	"fmt"
	"net/http"
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
	fmt.Println("Hello from Middleware");
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64){
			fmt.Println("Hello from Middleware");
			c.AbortWithStatus(http.StatusUnauthorized);
		}
		var user models.User
		config.DB.First(&user,claims["Issuer"]);
		fmt.Println("Hello from Middleware");	
		if( user.ID ==0){
			fmt.Println("Hello from Middleware");
			c.AbortWithStatus(http.StatusUnauthorized);

		}
		c.Set("user",user);
		fmt.Println("Hello from Middleware");
		c.Next();
	} else {
		fmt.Println("NO Authorization");
		fmt.Println(err)
	}




}