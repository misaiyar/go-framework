package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
    _ "github.com/go-sql-driver/mysql"
)

func init(){
    orm.RegisterModel(new(SystemUser),new(SystemModule),new(UserRole),new(UserPermission),new(SystemConfig))
    orm.RegisterDriver("mysql",orm.DRMySQL)
    //orm.Debug = true
    dn := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") +"@tcp(" + beego.AppConfig.String("mysqlurls")+ ")/" + beego.AppConfig.String("mysqldb") + "?charset=" + beego.AppConfig.DefaultString("mysqlchartset","utf8");
    orm.RegisterDataBase("default","mysql",dn)
}

//系统用户
type SystemUser struct {
    UserId int32 `orm:"column(user_id);pk;auto"`
    UserName string `orm:"column(username)"`
    Password string `orm:"column(password);size(32)"`
    Email string `orm:"column(email);unique"`
    RealName string `orm:"column(realname)"`
    RoleId int32 `orm:"column(role_id)"`
    Status int8 `orm:"column(status)"`
    Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func (u *SystemUser) TableName() string{
    return "system_user"
}

//系统模块
type SystemModule struct {
    ModuleId int32 `orm:"column(module_id);pk;auto"`
    ModuleName string `orm:"column(module_name)"`
    MethodName string `orm:"column(method_name)"`
    Title string `orm:"column(title)"`
    ParentId int32 `orm:"column(parent_id)"`
    Sort int32 `orm:"column(sort)"`
    Show int8 `orm:"column(show)"`
    Type int8 `orm:"column(type)"`
    Status int8 `orm:"column(status)"`
    Children map[int32]SystemModule `orm:"-"`
}

func (u *SystemModule) TableName() string{
    return "system_module"
}

//系统角色
type UserRole struct {
    RoleId int32 `orm:"column(role_id);pk;auto"`
    RoleName string `orm:"column(role_name)"`
    Description string `orm:"column(description)"`
    Count int64 `orm:"-"`
    Status int8 `orm:"column(status)"`
}

func (u *UserRole) TableName() string{
    return "system_role"
}

//用户权限
type UserPermission struct {
    Id int32 `orm:"column(id);pk"`
    RoleId int32 `orm:"column(role_id)"`
    ModuleId int32 `orm:"column(module_id)"`
}

func (u *UserPermission) TableName() string{
    return "system_access"
}

func (u *UserPermission) TableUnique() [][]string {
    return [][]string{
        []string{"RoleId", "ModuleId"},
    }
}
//系统设置
type SystemConfig struct {
    Key string `orm:"column(k);size(32);pk"`
    Value string `orm:"column(v);type(text)"`
    Description string `orm:"column(d);size(100)"`
}

type DicValue struct {
    Key string `json:"k"`
    Description string `json:"d"`
    Time int64 `json:"at"`
    Operate string `json:"opt"`
}

func (u *SystemConfig) TableName() string{
    return "system_setting"
}

//分页对象

type Pagination struct {
    CurrentPage int32 `json:"currentPage"`
    TotalCounts int32 `json:"totalCounts"`
    BaseUrl string `json:"baseUrl"`
}