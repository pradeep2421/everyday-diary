package routers

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/diary")
	{
		grp2 := grp1.Group("/user")
		{
			grp2.GET("/profile?id", controllers.GetUserProfile)
			grp2.POST("/register", controllers.RegisterUser)
			grp2.POST("/login", controllers.LoginUser)
			grp2.PATCH("/editProfile", controllers.EditUserProfile)
		}
		grp3 := grp1.Group("/diaryPage")
		{
			grp3.GET("/viewPage?id=1", controllers.GetDiaryPage)
			grp3.GET("/viewAllPage?", controllers.GetAllDiaryPage)
			grp3.DELETE("/delete?id", controllers.DeleteDiaryPage)
			grp3.POST("/add", controllers.AddDiaryPage)
			grp3.PATCH("/edit", controllers.EditDiaryPage)
		}
		grp4 := grp1.Group("/template")
		{
			grp4.GET("/list", controllers.GetAllTemplates)
			grp4.POST("/add", controllers.AddTemplate)
			grp4.PATCH("/edit", controllers.EditTemplate)
			grp4.DELETE("/delete", controllers.DeleteTemplate)
		}
		grp5 := grp1.Group("/posts")
		{
			grp5.POST("/add", controllers.AddPost)
			grp5.PATCH("/edit", controllers.EditPost)
			grp5.DELETE("/delete", controllers.DeletePost)
			grp5.GET("/post", controllers.GetPost)
			grp5.GET("/list", controllers.GetAllPost)
		}
		grp6 := grp1.Group("/comments")
		{
			grp6.POST("/add", controllers.AddComment)
			grp6.PATCH("/edit", controllers.EditComment)
			grp6.DELETE("/delete", controllers.DeleteComment)
			grp6.GET("/userComments", controllers.GetUserComment)
			grp6.GET("/postComments", controllers.GetPostComment)
		}
		grp7 := grp1.Group("/dashboard")
		{
			grp7.GET("/add", controllers.AddMonthData)
			grp7.GET("/month", controllers.GetDashBoard)
		}
	}
	return r
}