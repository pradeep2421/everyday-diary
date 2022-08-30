package Routes

import (
	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/diary")
	{
		grp1.GET("user",)
	}
	return r
}