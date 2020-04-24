package cust

import (
	"dc/models"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type controller struct {
	Controller string
	Action     string
}

func CasbinInsert(roleVo *models.AdminRole) {
	o := orm.NewOrm()
	//如果角色存在先将角色的权限全部删除
	models.AuthEnforcer.RemoveFilteredPolicy(0, strconv.Itoa(int(roleVo.Rid)))

	//插入角色权限
	roleVoMidArray := strings.Split(roleVo.Mids, ",")
	menus := []models.AdminMenu{}
	_, e := o.QueryTable(new(models.AdminMenu).TableName()).Filter("mid__in", roleVoMidArray).All(&menus)
	if nil != e {
		log.Println("Casbib角色权限插入失败:", e)
	}

	//将controller和action取出 并 去重
	controllerSlice := []controller{}

	for _, menu := range menus {
		c := controller{}
		c.Controller = menu.Controller
		c.Action = menu.Action
		repeat := false
		for _, v := range controllerSlice {
			if v.Controller == menu.Controller && v.Action == menu.Action {
				repeat = true
			}
		}
		//去重
		if !repeat {
			controllerSlice = append(controllerSlice, c)
		}
	}

	//去掉切片第一个的元素的空值
	controllerSlice = append(controllerSlice[1:])

	//更新角色权限
	for _, v := range controllerSlice {
		models.AuthEnforcer.AddNamedPolicy("p", strconv.Itoa(int(roleVo.Rid)), v.Controller, v.Action)
	}
}
