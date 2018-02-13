package controllers

import (
    "project/models"
    "crypto/md5"
    "encoding/hex"
)

type SysuserController struct {
    baseAdminController
}

func (c *SysuserController) Index() {
    page,_ := c.GetInt32("page",1)
    user_name := c.GetString("uname","")
    role_lists,_ := new(models.UserRole).GetRoleList()
    c.Data["RoleList"] = &role_lists
    user_list,pagination := new(models.SystemUser).GetPage(page,10,"/Sysuser/Index",user_name)
    c.Data["UserList"] = &user_list
    c.Data["Pagination"] = &pagination
    c.Layout ="common/layout.html"
    c.TplName = "user/index.html"
}

func (c *SysuserController) Save() {
    var user_name = c.GetString("username")
    var password = c.GetString("password")
    var pwdcheck = c.GetString("pwdcheck")
    var email = c.GetString("email")
    var realname = c.GetString("realname")
    var role_id,_ = c.GetInt32("role_id",int32(0))
    var status,_ = c.GetInt8("status",int8(0))
    var uid,_ = c.GetInt32("user_id",0)
    if uid==0 && user_name=="" || email==""|| realname==""|| role_id==0{
        c.RenderJson(300,"缺少必要参数",nil)
    }
    h := md5.New()
    h.Write([]byte(password)) // 需要加密的字符串
    password_md5 := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
    h = md5.New()
    h.Write([]byte(pwdcheck)) // 需要加密的字符串
    pwdcheck_md5 := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
    if password_md5 != pwdcheck_md5{
        c.RenderJson(300,"密码有误",nil)
    }
    user := models.SystemUser{UserId:uid,UserName:user_name,Password:password_md5,
        Email:email,RoleId:role_id,Status:status,RealName:realname}
    if ok,err := user.Save(); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}

func (c *SysuserController) UpdateStatus() {
    var status,_ = c.GetInt8("status",int8(0))
    var uid,_ = c.GetInt32("user_id",0)
    if uid==0 {
        c.RenderJson(300,"缺少必要参数",nil)
    }
    user := models.SystemUser{UserId:uid,Status:status}
    if ok,err := user.Update( "Status" ); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}

func (c *SysuserController) ResetPasswd() {
    var uid,_ = c.GetInt32("user_id",0)
    var password = c.GetString("password")
    var pwdcheck = c.GetString("pwdcheck")
    if uid==0 || password==""|| pwdcheck=="" {
        c.RenderJson(300,"缺少必要参数",nil)
    }
    h := md5.New()
    h.Write([]byte(password)) // 需要加密的字符串
    password_md5 := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
    h = md5.New()
    h.Write([]byte(pwdcheck)) // 需要加密的字符串
    pwdcheck_md5 := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
    if password_md5 != pwdcheck_md5{
        c.RenderJson(300,"密码有误",nil)
    }
    user := models.SystemUser{UserId:uid,Password:password_md5}
    if ok,err := user.Update( "Password" ); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}
