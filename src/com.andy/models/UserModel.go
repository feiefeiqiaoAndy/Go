package models

type UserModel struct {
    Uid   int
	Uname string
}

func (u UserModel) SetValue (id int ,name string)   {

	u.Uid = id
	u.Uname = name
}


func (u UserModel)    () string {

	return  "用户名是"+u.Uname
}
