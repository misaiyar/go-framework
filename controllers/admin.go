package controllers

type AdminController struct {
    baseAdminController
}

func (c *AdminController) Main() {
    c.Layout ="common/layout.html"
    c.TplName = "main.html"
}
//退出系统
func (c *AdminController) Logout(){
    c.DestroySession()
    c.Redirect("/",302)
    return
}
