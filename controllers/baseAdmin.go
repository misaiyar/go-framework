package controllers

import (
	"github.com/astaxie/beego"
	"project/models"
	"fmt"
)

type baseAdminController struct {
	BaseController
}

func (c *baseAdminController) Prepare() {
	//输出layout需要的数据
	c.Data["SystemName"] = beego.AppConfig.String("systenName")
	c.Data["SystemAlias"] = beego.AppConfig.String("systenAlias")
	//判断登录
	if c.GetSession("user_id") == nil {
		c.Redirect("/",302)
		return
	}
	c.Data["LoginUserName"] = c.GetSession("username")
	c.Data["LoginUserEmail"] = c.GetSession("email")
	control,action:= c.GetControllerAndAction()
	control = beego.Substr(control,0,len(control)-10)
	var role_id int32= 0
	session_role_id := c.GetSession("role_id")
	if session_role_id !=nil {
		role_id = session_role_id.(int32)
	}
	//判断url权限
	access_tmp := c.GetSession("UserPermission");
	access := map[string]map[string]bool{"Admin":map[string]bool{"Main":true,"Logout":true}}//设置默认页面
	if access_tmp ==nil || len(access)==0 {
		tmp_access := access
		access = new(models.SystemModule).GetAccessModuleList(role_id)
		if len(access)==0 {
			access = tmp_access
		}
		c.SetSession("UserPermission",access)
	}else{
		access = access_tmp.(map[string]map[string]bool)
	}
	fmt.Print("Permission Access:"+control + ">>>" + action + ":")
	v,_ := access[control][action]
	fmt.Println(v)
	if v,ok := access[control][action]; !( ok && v ) {
		if( c.IsAjax() ){
			c.RenderJson(400,"您没有操作权限",nil)
		}else{
			c.RenderScript("您没有操作权限",true)
		}
	}
	roleMenus_tmp := c.GetSession("UserMenu");
	roleMenus := make(map[int32]models.ParentMenu)
	if roleMenus_tmp ==nil {
		roleMenus, _ = new(models.UserRole).GetMenus(role_id)
		c.SetSession("UserMenu",roleMenus)
	}else{
		roleMenus = roleMenus_tmp.(map[int32]models.ParentMenu)
	}
	c.Data["UserMenu"]=&roleMenus
	c.Data["Module"]= control
	c.Data["Action"]= action
}