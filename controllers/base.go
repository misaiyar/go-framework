package controllers

import "github.com/astaxie/beego"

type BaseController struct {
    beego.Controller
}

type ResponseJson struct {
    Code int32 `json:"code"`
    Data interface{} `json:"data"`
    Message string `json:"msg"`
}

func (c *BaseController) RenderJson(code int32,message string,data interface{}) {
    mystruct := ResponseJson{Code:code,Data:data,Message:message}
    c.Data["json"] = &mystruct
    c.ServeJSON()
    c.StopRun()
}

func (c *BaseController) RenderScript(message string,back bool)  {
    ext_str := ""
    if back {
        ext_str += "window.history.back();"
    }
    content := "<script>alert(\""+ message +"\");"+ ext_str +"</script>"
    c.Ctx.Output.Body([]byte(content))
    c.StopRun()
}