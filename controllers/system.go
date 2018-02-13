package controllers

import (
    "project/models"
    "time"
)

type SystemController struct {
    baseAdminController
}

func (c *SystemController) Index() {
    var key = c.GetString("type","")
    var current = new(models.SystemConfig)
    dictList,divValues := current.GetList(key)
    c.Data["List"] = &divValues
    c.Data["DictList"] = &dictList
    if( current.Key == "" ){
        current.Description = "请先添加字典"
    }
    c.Data["Current"] = &current
    c.Layout ="common/layout.html"
    c.TplName = "system/dictlist.html"
}

func (c *SystemController) Adddict(){
    var key = c.GetString("k","")
    var desc = c.GetString("d","")
    mystruct := ResponseJson{200,nil,"处理成功"}
    if key=="" || desc =="" {
        mystruct.Message = "缺少必要参数"
        mystruct.Code = 300
        c.Data["json"] = &mystruct
        c.ServeJSON()
        return
    }
    config := models.SystemConfig{Key:key,Description:desc}
    if ok,err := config.Save(nil,nil); !ok {
        mystruct.Message = err.Error()
        mystruct.Code = 300
    }
    c.Data["json"] = &mystruct
    c.ServeJSON()
}

func (c *SystemController) Savedictvalue(){
    var k_type = c.GetString("type","")
    var key = c.GetString("k","")
    var new_key = c.GetString("newk","")
    var desc = c.GetString("d","")
    if key=="" && new_key=="" || desc =="" {
        c.RenderJson(300,"缺少必要参数",nil)
    }
    username := c.GetSession("username").(string)
    config := models.SystemConfig{Key:k_type,Description:desc}
    var add_dic_valus = []models.DicValue{{Key:new_key,Description:desc,
        Operate:username,Time:time.Now().Unix()}}
    if ok,err := config.Save(add_dic_valus,[]models.DicValue{{Key:key}}); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}

func (c *SystemController) Deldictvalue(){
    var k_type = c.GetString("type","")
    var keys = c.GetStrings("ids[]")
    if k_type=="" || keys == nil {
        c.RenderJson(300,"缺少必要参数",nil)
    }
    var del_dic_values []models.DicValue;
    for _, value := range keys {
        del_dic_values = append(del_dic_values,models.DicValue{Key:value})
    }
    config := models.SystemConfig{Key:k_type}
    if ok,err := config.Save(nil,del_dic_values); !ok {
        c.RenderJson(300,err.Error(),nil)
    }
    c.RenderJson(200,"处理成功",nil)
}
