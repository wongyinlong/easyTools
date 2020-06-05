package models

import "github.com/astaxie/beego/orm"

type AdminRole struct {
	Rid    int64  `orm:"column(rid);pk;auto"json:"rid"`
	RName  string `orm:"column(r_name);size(50)"json:"r_name"`
	RCode  string `orm:"column(r_code);size(50)"json:"r_code"`
	Mids   string `orm:"column(mids);size(2000)"json:"mids"`
	Desc   string `orm:"column(desc);size(50)"json:"desc"`
	Remark string `orm:"column(remark);size(300)"json:"remark"`
}

func (r *AdminRole) TableName() string {
	return "admin_role"
}

func RegisterAdminRole() {
	orm.RegisterModel(new(AdminRole))
}
