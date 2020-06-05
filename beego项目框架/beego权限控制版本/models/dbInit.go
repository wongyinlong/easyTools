package models

import (
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	beegoormadapter "github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// 引入了casbin
	AuthEnforcer *casbin.Enforcer
)

func init() {
	User := beego.AppConfig.String("mysql.user")
	Pwd := beego.AppConfig.String("mysql.pwd")
	Host := beego.AppConfig.String("mysql.host")
	DbName := beego.AppConfig.String("mysql.dbname")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", User, Pwd, Host, DbName)
	maxIdle := 30
	maxConn := 30

	orm.Debug = beego.AppConfig.DefaultBool("mysql.debug", false)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("数据库连接异常!")
		os.Exit(-1)
	}

	err = orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
	if err != nil {
		logs.Error("数据库连接异常, 请检查服务器!")
		os.Exit(-1)
	}
	// 注册数据库
	RegisterAdminRole()
	RegisterAdminUser()
	RegisterExampleModel()

	//初始化权限控制
	enforcerAuth(dataSource) // 初始化权限控制
	//这里还需 user表 和role表。 在权限控制查表时用
	orm.RunSyncdb("default", false, true)
}

func enforcerAuth(dataSource string) {
	var err error

	casbinConf := beego.AppConfig.DefaultString("casbin.model", "conf/casbin.conf")

	AuthEnforcer, err = casbin.NewEnforcerSafe(casbinConf, beegoormadapter.NewAdapter("mysql", dataSource, true))
	if nil != err {
		logs.Error(err.Error())
		return
	}

	if err := AuthEnforcer.LoadPolicy(); err != nil {
		logs.Error(err.Error())
		return
	}
}

// 插入对象
func Insert(entity interface{}) int64 {
	num, err := orm.NewOrm().InsertOrUpdate(entity)
	if err != nil {
		log.Println("对象：", entity, "数据插入失败：", err)
		return 0
	}

	return num
}

// 删除对象
func Delete(item interface{}) error {
	if _, err := orm.NewOrm().Delete(item); err != nil {
		log.Println("对象：", item, "数据删除失败：", err)
		return err
	} else {
		return nil
	}
}

// 查询Entity详情
func Query(item interface{}, cols ...string) error {
	return orm.NewOrm().Read(item, cols...)
}

// 按照ID修改
func Update(item interface{}, cols ...string) (int64, error) {
	return orm.NewOrm().Update(item, cols...)
}
