package services

import (
	"backend/config"
	"backend/models"
	"fmt"
)

func CreateNote(note *models.Note) (err error) {
	err = config.DB.Create(&note).Error;

	return err;
}
func UpdateNote(note *models.Note, id string) (newNote models.Note,err error){
	
	err = config.DB.First(&newNote,id).Error;
	

	fmt.Println(note);
	config.DB.Model(&newNote).Updates(models.Note{
		Name: note.Name,
		Description: note.Description,
	})
	
	return newNote ,err;
}

func DeleteNote(id string) (err error) {
	
	err = config.DB.Delete(&models.Note{},id).Error;

	return err;

	
}
func ViewNote(notes *[]models.Note) (err error) {
	
err = config.DB.Find(&notes).Error;

	return err;
}
func SingleNote(note *models.Note,id string) (err error) {
	
	err = config.DB.Find(&note,id).Error;
	return err;
}