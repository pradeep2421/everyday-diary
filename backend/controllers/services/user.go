package services

import (
	"backend/config"
	"backend/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) (err error) {

	err = config.DB.Create(&user).Error;
	return err
}
func LoginUser(login *models.Login) (user models.User,err error) {
	config.DB.Where("name = ?",login.Email).First(&user);
	
	if(user.ID ==0){
		err = errors.New("user not Found");
		
		return user, err;
	}
	
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login.Password)); err !=nil {
		return user ,err;
	}
	
	return user, err;
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