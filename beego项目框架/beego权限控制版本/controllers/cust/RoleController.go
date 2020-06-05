package cust

import (
	"dc/controllers/basic"
	"dc/dao"
	"dc/models"
	"dc/utils"
	"dc/vo"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type RoleController struct {
	basic.AuthorizeController
}

func (c *RoleController) Router() *beego.Namespace {
	ns := beego.NewNamespace("/role",
		beego.NSRouter("/list", c, "get:List"),
		beego.NSRouter("/add", c, "post:Add"),
		beego.NSRouter("/info/:rid", c, "get:Info"),
		beego.NSRouter("/update/:rid", c, "put:Update"),
		beego.NSRouter("/delete/:rid", c, "delete:Delete"),
	)

	return ns
}

func (c *RoleController) List() {
	pageNo, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("limit", 20)

	page := models.NewPage(pageNo-1, pageSize)

	e := models.AdminRole{}
	// 1. 根据条件查询数据分页
	r := dao.FindRolePage(page, &e)
	c.JSONOk(r)
	c.StopRun()
}

func (c *RoleController) Add() {
	r := map[string]interface{}{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &r)
	if nil != err {
		c.JSONErr(vo.Error.SetMsg("接受的数据不是json格式"))
		c.StopRun()
	}

	//自动生成角色代码
	rCodeLength, err := beego.AppConfig.Int("role.r_code.length")
	if nil != err {
		rCodeLength = 6 //默认6位字符串
	}
	rCode := utils.RndUpperString(rCodeLength)

	//数据入库
	var rName, remark, mids, desc string
	if v, ok := r["r_name"].(string); ok {
		rName = v
	}
	if v, ok := r["remark"].(string); ok {
		remark = v
	}
	if v, ok := r["mids"].(string); ok {
		mids = v
	}
	if v, ok := r["desc"].(string); ok {
		desc = v
	}

	adminRole := models.AdminRole{
		RCode:  rCode,
		RName:  rName,
		Remark: remark,
		Mids:   mids,
		Desc:   desc,
	}

	rid := models.Insert(&adminRole)
	if rid != 0 {
		c.JSONOk(map[string]interface{}{"role": adminRole})

		//新建角色成功后加入casbin表
		adminRole.Rid = rid
		CasbinInsert(&adminRole)

	}
	c.JSONErr(vo.ErrInputData)
	c.StopRun()
}

func (c *RoleController) Info() {
	ridStr := c.Ctx.Input.Param(":rid")

	ridInt, err := strconv.Atoi(ridStr)
	if (nil == err && ridInt <= 0) || nil != err {
		c.JSONErr(vo.ErrInputData.SetMsg("请求的数据不正确"))
		c.StopRun()
	}

	adminRole := models.AdminRole{Rid: int64(ridInt)}
	err = models.Query(&adminRole)
	if nil != err {
		c.JSONErr(vo.ErrNoUser.SetMsg("查询失败"))
		c.StopRun()
	}

	midsStr := strings.Split(adminRole.Mids, ",")
	var midsInt []int64
	for _, midStr := range midsStr {
		v, _ := strconv.Atoi(midStr)
		midsInt = append(midsInt, int64(v))
	}

	menus := dao.FindAdminMenuListByMids(midsInt)

	c.JSONOk(map[string]interface{}{"role": adminRole, "menus": menus})
	c.StopRun()
}

func (c *RoleController) Update() {
	ridStr := c.Ctx.Input.Param(":rid")
	rid, err := strconv.Atoi(ridStr)
	if (nil == err && rid <= 0) || nil != err {
		c.JSONErr(vo.ErrInputData.SetMsg("请求的数据不正确"))
		c.StopRun()
	}
	r := map[string]interface{}{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &r)
	if nil != err {
		c.JSONErr(vo.Error.SetMsg("接受的数据不是json格式"))
		c.StopRun()
	}

	//数据更新入库
	var rName, remark, mids, desc string
	if v, ok := r["r_name"].(string); ok {
		rName = v
	}
	if v, ok := r["remark"].(string); ok {
		remark = v
	}
	if v, ok := r["mids"].(string); ok {
		mids = v
	}
	if v, ok := r["desc"].(string); ok {
		desc = v
	}

	adminRoleVo := models.AdminRole{
		Rid:    int64(rid),
		RName:  rName,
		Remark: remark,
		Mids:   mids,
		Desc:   desc,
	}

	updateKey := []string{"r_name", "mids", "desc", "remark"}

	o := orm.NewOrm()
	_, err = o.Update(&adminRoleVo, updateKey...)
	if nil != err {
		c.JSONErr(vo.ErrInputData.SetMsg("更新失败"))
		c.StopRun()
	}

	//更新casbin权限表
	CasbinInsert(&adminRoleVo)

	c.JSONOk()
	c.StopRun()
}

func (c *RoleController) Delete() {
	ridStr := c.Ctx.Input.Param(":rid")
	rid, err := strconv.Atoi(ridStr)
	if (nil == err && rid <= 0) || nil != err {
		c.JSONErr(vo.ErrInputData.SetMsg("请求的数据不正确"))
		c.StopRun()
	}

	adminRole := models.AdminRole{Rid: int64(rid)}
	err = models.Delete(&adminRole)
	if nil != err {
		c.JSONErr(vo.ErrNoUserChange)
		c.StopRun()
	}

	c.JSONOk()
	c.StopRun()

}
