package services

import (
	"backend/config"
	"backend/models"
	"fmt"
)

func CreateUser(user *models.User) (err error) {

	err = config.DB.Create(&user).Error;
	return err
}
func LoginUser(login *models.Login) (user models.User,err error) {

	config.DB.Where(&models.User{Name: login.Name, Password: login.Password}).First(&user)
	fmt.Println(user);
	return user, err
}

func UpdateUser(user *models.User) (err error){
	

	config.DB.Save(&user);
	
	return err;
}

func DeleteUser(id string) (err error) {
	
	err = config.DB.Delete(&models.User{},id).Error;

	return err;

	
}
func ViewUser(users *[]models.User) (err error) {
	
err = config.DB.Find(&users).Error;

	return err;
}
func SingleUser(user *models.User,id string) (err error) {
	
	err = config.DB.Find(&user,id).Error;
	return err;
}