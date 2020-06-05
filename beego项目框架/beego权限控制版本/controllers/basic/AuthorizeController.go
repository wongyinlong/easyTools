package basic

import (
	"dc/basic"
	"dc/models"
	"dc/vo"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

/**
 * 控制器授权Controller基础
 */
type AuthorizeController struct {
	BaseController
}

func (c *AuthorizeController) Prepare() {
	c.BaseController.Prepare()

	auth := c.Ctx.Input.Header(c.TokenHeader())

	if auth == "" {
		c.JSONErr(vo.ErrExpired.SetMsg("TOKEN不存在"))
	} else {

		// todo 这里需要判断缓存是否在缓存中存在

		claims, ok := basic.ParseToken(auth, string(basic.EncryptKey))
		if !ok {
			c.JSONErr(vo.ErrExpired.SetMsg("TOKEN无效或者过期"))
		} else {
			c.SetSession("uid", claims.(jwt.MapClaims)["jti"].(string))
		}
	}

	//2.权限验证

	//角色退出，获取当前用户角色和获取菜单资源不需要权限验证
	controllerName, activeName := c.ControllerName, c.ActionName

	if !verify(controllerName) {
		return
	}

	currUserRoleID := c.GetCurrUserRoleID()

	policy := models.AuthEnforcer.GetFilteredPolicy(0, strconv.Itoa(int(currUserRoleID)))
	if len(policy) <= 0 {
		c.JSONErr(vo.ErrPermission)
		c.StopRun()
	} else {

		action := setAction(activeName)

		result, err := models.AuthEnforcer.EnforceSafe(strconv.Itoa(int(currUserRoleID)), controllerName, action)
		if nil != err {
			c.JSONErr(vo.ErrPermission)
			c.StopRun()
		}
		if !result {
			c.JSONErr(vo.ErrPermission)
			c.StopRun()
		}
	}
}

func setAction(action string) string {

	if action == "List" || action == "Info" || action == "LogoUpload" || action == "AllList" || action == "SortInfo" || action == "VestList" || action == "GetSet" || action == "ChannelCnt" || action == "AgreeInfo" {
		action = "List"
	}

	if action == "SortUpdate" {
		action = "Update"
	}

	if action == "Set" {
		action = "Add"
	}
	return action
}

func verify(controllerName string) bool {
	if controllerName == "UserController" {
		return false
	}
	return true
}
