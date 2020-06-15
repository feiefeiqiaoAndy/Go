package services

import "com.andy/models"

func  GetUser() string {

    user :=new(models.UserModel)

     user.SetValue(11,"qiao")



	return   user.ToString()

}