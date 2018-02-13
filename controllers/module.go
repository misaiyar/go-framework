package controllers

import "project/models"

type ModuleController struct {
    baseAdminController
}

func (c *ModuleController) Index() {
    module_lists,_ := new(models.SystemModule).GetModuleList()
    c.Data["ModuleData"] =&module_lists
    c.Layout ="common/layout.html"
    c.TplName = "module/index.html"
}

func (c *ModuleController) Save() {
    var title = c.GetString("title")
    var parent_id,err = c.GetInt32("parent_id")
    var module_name = c.GetString("module_name")
    var method_name = c.GetString("method_name")
    var types,_ = c.GetInt8("type",int8(0))
    var show,_= c.GetInt8("show",int8(0))
    var sort,_= c.GetInt32("sort",0)
    var status,_ = c.GetInt8("status",int8(0))
    var id,_ = c.GetInt32("id",0)
    if title==""|| err!=nil || module_name==""{
        c.RenderJson(300,"缺少必要参数",nil)
    }
    module := models.SystemModule{ModuleId:id,ModuleName:module_name,MethodName:method_name,
        Title:title,ParentId:parent_id,Sort:sort,Show:show,Type:types,Status:status,Children:nil}
    if ok,err := module.Save(); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}


