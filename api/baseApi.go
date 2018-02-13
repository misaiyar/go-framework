package api

import (
	. "project/controllers"
	"project/libs"
)

type baseApiController struct {
	BaseController
}

func (c *baseApiController) Prepare() {
	api_key := c.GetString("apikey")
	sign := c.GetString("sign")
	if api_key == "" || sign=="" {
		c.RenderJson(400,"缺少必要验签参数",nil)
	}
	form_data := c.Ctx.Request.Form
	var param = make(map[string]string)
	for k,v := range form_data {
		if k == "sign" {
			continue
		}
		param[k]= v[len(v)-1]
	}
	var secret_key,ok = libs.AuthApp[api_key]
	if !ok {
		c.RenderJson(400,"apikey无效",nil)
	}
	if sign != libs.GenToken(secret_key,param) {
		c.RenderJson(400,"签名不正确",nil)
	}
}