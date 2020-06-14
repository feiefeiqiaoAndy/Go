package services

import "com.andy/models"

func  GetUser() string {

    user :=new(models.UserModel)
    user.Uname = "andy"

	return   user.ToString()

}