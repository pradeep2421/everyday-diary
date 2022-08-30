package Routers

import (
	"backend/Controllers"

	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/diary")
	{
		grp2 := grp1.Group("/user")
		{
			grp2.GET("/profile?id",Controllers.GetUserProfile)
			grp2.POST("/register", Controllers.RegisterUser)
			grp2.POST("/login",Controllers.LoginUser)
			grp2.PATCH("/editProfile",Controllers.EditUserProfile)
		}
		grp3 := grp1.Group("/diaryPage")
		{
			grp3.GET("/viewPage?id=1",Controllers.GetDiaryPage)
			grp3.GET("/viewAllPage?",Controllers.GetAllDiaryPage)
			grp3.DELETE("/delete?id",Controllers.DeleteDiaryPage)
			grp3.POST("/add",Controllers.AddDiaryPage)
			grp3.PATCH("/edit",Controllers.EditDiaryPage)
		}
		grp4 := grp1.Group("/template")
		{
			grp4.GET("/list",Controllers.GetAllTemplates)
			grp4.POST("/add",Controllers.AddTemplate)
			grp4.PATCH("/edit",Controllers.EditTemplate)
			grp4.DELETE("/delete",Controllers.DeleteTemplate)
		}
		grp5 := grp1.Group("/posts")
		{
			grp5.POST("/add",Controllers.AddPost)
			grp5.PATCH("/edit",Controllers.EditPost)
			grp5.DELETE("/delete",Controllers.DeletePost)
			grp5.GET("/post",Controllers.GetPost)
			grp5.GET("/list",Controllers.GetAllPost)
		}
		grp6 := grp1.Group("/comments")
		{
			grp6.POST("/add",Controllers.AddComment)
			grp6.PATCH("/edit",Controllers.EditComment)
			grp6.DELETE("/delete",Controllers.DeleteComment)
			grp6.GET("/userComments",Controllers.GetUserComment)
			grp6.GET("/postComments",Controllers.GetPostComment)
		}
		grp7 := grp1.Group("/dashboard")
		{
			grp7.GET("/add",Controllers.AddMonthData)
			grp7.GET("/month",Controllers.GetDashBoard)
		}
	}
	return r
}