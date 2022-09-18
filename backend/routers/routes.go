package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	grp1 := r.Group("/diary")
	{
		grp2 := grp1.Group("/user")
		{
			
			grp2.POST("/create", controllers.CreateUser)
			grp2.POST("/login", controllers.LoginUser)
			grp2.PATCH("/update/:id", controllers.UpdateUser)
			grp2.GET("/view", controllers.ViewUser)
			grp2.GET("/singleuser/:id",middleware.RequireAuth, controllers.SingleUser)
			grp2.DELETE("/delete/:id", controllers.DeleteUser)
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
		grp8 := grp1.Group("/note")
		{
			grp8.POST("/create", controllers.CreateNote)
			grp8.PATCH("/update/:id", controllers.UpdateNote)
			grp8.DELETE("/delete/:id", controllers.DeleteNote)
			grp8.GET("/view", controllers.ViewNote)
			grp8.GET("/singlenote/:id",controllers.SingleNote)
		}
	}
	return r
}