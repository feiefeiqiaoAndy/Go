package models

type UserModel struct {
	Uid int
	Uname string




}

func (u UserModel) ToString () string {

	return  "用户名是"+u.Uname
}
