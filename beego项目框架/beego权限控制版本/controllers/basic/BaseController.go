package basic

import (
	"dc/models"
	"dc/utils"
	. "dc/vo"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

/**
 *  基础Controller封装
 */
type BaseController struct {
	beego.Controller
	ControllerName string //当前控制名称
	ActionName     string //当前action名称
}

func (c *BaseController) Prepare() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Origin,X-Requested-With, Content-Type, Accept, accept-encoding, Authorization")
	c.Ctx.Output.Header("Access-Control-Max-Age", "1728000")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	//设置控制器和动作信息
	c.ControllerName, c.ActionName = c.GetControllerAndAction()
}

func (c *BaseController) TokenHeader() string {
	return beego.AppConfig.String("token.header")
}

// 获取真实的IP地址

func (c *BaseController) RemoteIp() string {

	ip := utils.RemoteIp(c.Ctx.Request)
	realIp := ""
	if strings.Contains(ip, ":") {
		ips := strings.Split(ip, ":")
		realIp = ips[0]
	} else {
		realIp = ip
	}

	return realIp
}

func (c *BaseController) JSON(status int, d interface{}) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = d
	c.ServeJSON(true)
}

// 返回错误结果
func (c *BaseController) JSONErr(d interface{}) {
	resp, ok := d.(*Result)
	if ok {
		c.JSON(200, resp)
	} else {
		c.JSON(200, Error.SetData(d))
	}
}

// 返回正确的结果
func (c *BaseController) JSONOk(r ...interface{}) {
	if len(r) <= 0 {
		c.JSON(200, Success)
		return
	}

	resp, ok := r[0].(*Result)
	if ok {
		c.JSON(200, resp)
	} else {
		c.JSON(200, Success.SetData(r))
	}
}

//获取当前用户ID
func (c *BaseController) GetCurrUserID() int64 {
	session := c.GetSession("uid")
	if nil == session {
		c.JSONErr(ErrExpired)
		c.StopRun()
	}

	var uid int
	if v, ok := session.(string); ok {
		uid, _ = strconv.Atoi(v)
	}

	return int64(uid)
}

//获取当前用户角色ID
func (c *BaseController) GetCurrUserRoleID() int64 {

	uid := c.GetCurrUserID()

	adminUser := models.AdminUser{Uid: int64(uid)}
	if nil != models.Query(&adminUser) {
		c.JSONErr(ErrExpired)
		c.StopRun()
	}

	return adminUser.Rid
}

//获取当前用户角色名称
func (c *BaseController) GetCurrRoleName() string {

	uid := c.GetCurrUserID()

	adminUser := models.AdminUser{Uid: int64(uid)}
	if nil != models.Query(&adminUser) {
		c.JSONErr(ErrExpired)
		c.StopRun()
	}

	adminRole := models.AdminRole{Rid: adminUser.Rid}
	if nil != models.Query(&adminRole) {
		c.JSONErr(ErrExpired)
		c.StopRun()
	}

	return adminRole.RName
}

// 重定向 去错误页
func (c *BaseController) PageError() {
	errURL := c.URLFor("LoginController.PageError404")
	c.Redirect(errURL, 302)
}

func DataAuth(r RMap, data, typ string) (bool, interface{}) {

	switch typ {
	case "float64":
		if v, ok := r[data].(float64); ok {
			if v >= 0 {
				return true, v
			}
		}
		return false, 0
	case "int64":
		if v, ok := r[data].(int64); ok {
			if v >= 0 {
				return true, v
			}
		}
		return false, 0
	default:
		if v, ok := r[data].(string); ok {
			if len(v) >= 0 {
				return true, v
			}
		}
		return false, ""
	}

}
