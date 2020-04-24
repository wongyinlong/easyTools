package models

import "github.com/astaxie/beego/orm"

// 渠道
type Example struct {
	ID    int64  `orm:"column(cid);pk;auto" json:"cid"`        //ID
	UID   int64  `orm:"column(uid)" json:"uid"`                // UID
	VID   int64  `orm:"column(vid)" json:"vid"`                // VID
	CName string `orm:"column(c_name);size(30)" json:"c_name"` // CName

}

func (c *Channel) TableName() string {
	return "Example"
}

func RegisterExampleModel() {
	orm.RegisterModel(new(Channel), new(ChannelPv))
}
