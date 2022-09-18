package controllers

import (
	"backend/controllers/services"
	"backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)


var err error;
func CreateNote(c *gin.Context) {
	var note models.Note;
	c.BindJSON(&note);
	err = services.CreateNote(&note);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":"cannot create note"});
		return;
	}
	c.JSON(200, gin.H{"note":note});

	
}
func UpdateNote(c *gin.Context) {
	var note models.Note;
	id := c.Param("id");
	c.BindJSON(&note);
	fmt.Println(note);
	note,err := services.UpdateNote(&note, id);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":"unable to update note"});
		return;
	}
	c.JSON(200, gin.H{"updated_note":note});

}

func DeleteNote(c *gin.Context) {
	//get the id off the url
	id := c.Param("id");
	
	err := services.DeleteNote(id);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":"cannot delete note","error":err});
		return;
	}
	c.JSON(200,gin.H{"message":"note deleted"});
}
func ViewNote(c *gin.Context) {
	 var notes []models.Note;
	err = services.ViewNote(&notes);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":err});
		return;
	}
	c.JSON(200, gin.H{"notes":notes});
	
}
func SingleNote(c *gin.Context){
	var note models.Note;
	id := c.Param("id");
	err  = services.SingleNote(&note,id);
	if err != nil{
		c.AbortWithStatusJSON(400, gin.H{"message":err});
		return;
	}
	c.JSON(200, gin.H{"note":note});
}

