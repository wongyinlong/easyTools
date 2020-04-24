package models

import "github.com/astaxie/beego/orm"

type AdminUser struct {
	Uid      int64  `orm:"column(uid);pk;auto"json:"uid"`
	Rid      int64  `orm:"column(rid)"json:"rid"`
	UName    string `orm:"column(u_name)"json:"u_name"`
	UserName string `orm:"column(username);size(50)"json:"user_name"`
	Password string `orm:"column(password);size(100)"`
	Desc     string `orm:"column(desc);size(50)"json:"desc"`
	Pwd      string `orm:"column(pwd);size(50)"json:"pwd"`
	CreateAt int64  `orm:"column(create_at);type(bigint)"json:"create_at"`
	UpdateAt int64  `orm:"column(update_at);type(bigint)"json:"update_at"`
}

func (a *AdminUser) TableUnique() [][]string {

	return [][]string{
		[]string{"username"},
	}

}

func (a *AdminUser) TableName() string {
	return "admin_user"
}

func RegisterAdminUser() {
	orm.RegisterModel(new(AdminUser))
}
