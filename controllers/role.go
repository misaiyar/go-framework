package controllers

import (
    "project/models"
    "strconv"
)

type RoleController struct {
    baseAdminController
}

func (c *RoleController) Index() {
    role_lists,_ := new(models.UserRole).GetRoleList()
    c.Data["RoleData"] =&role_lists
    c.Layout ="common/layout.html"
    c.TplName = "role/index.html"
}

func (c *RoleController) Access() {
    var role_id,_ = c.GetInt32("role_id",0)
    access_lists := new(models.UserPermission).GetAccessList(role_id)
    c.Data["AccessData"] =&access_lists
    c.Data["RoleId"] =role_id
    module_lists,_ := new(models.SystemModule).GetModuleList()
    c.Data["ModuleData"] =&module_lists
    c.Layout ="common/layout.html"
    c.TplName = "role/access.html"
}

func (c *RoleController) SetAccess(){
    var item = c.GetStrings("item")
    var id,_ = c.GetInt32("role_id",0)
    if id==0 || item == nil {
        c.RenderJson(300,"缺少必要参数",nil)
    }
    var permissions []models.UserPermission;
    for _, value := range item {
        module_id,_ := strconv.ParseInt(value, 10, 32)
        permissions = append(permissions,models.UserPermission{RoleId:id,
            ModuleId:int32(module_id)})
    }
    if ok:= new(models.UserPermission).SetPermission(id,permissions); !ok {
        c.RenderJson(300,"设置失败",nil)
    }
    c.RenderJson(200,"处理成功",nil)
}

func (c *RoleController) ViewRole(){
    var id,_ = c.GetInt32("role_id",0)
    c.Data["RoleName"] = new(models.UserRole).GetRoleName(id)
    c.Data["UserList"] = new(models.SystemUser).GetRoleUsers(id)
    c.TplName = "role/viewRole.html"
}

func (c *RoleController) Save(){
    var name = c.GetString("role_name")
    var desc = c.GetString("description")
    var status,_ = c.GetInt8("status",int8(0))
    var id,_ = c.GetInt32("role_id",0)
    if name=="" {
        c.RenderJson(300,"缺少必要参数",nil)
    }
    module := models.UserRole{RoleId:id,RoleName:name,Description:desc,Status:status}
    if ok,err := module.Save(); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}
