package main

import (
	_ "project/routers"
	"github.com/astaxie/beego"
	"project/models"
	"encoding/gob"
)

func init()  {
	//session 值类型注册
	gob.Register(map[int32]models.ParentMenu{})
	gob.Register(map[string]map[string]bool{})
}
func main() {
	beego.Run()
}

