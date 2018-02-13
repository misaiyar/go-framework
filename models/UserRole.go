package models

import (
    "github.com/astaxie/beego/orm"
    "errors"
)

type Menu struct{
    Url string `json:"url"`
    Title string `json:"title"`
    Module string `json:"module_name"`
    Action string `json:"method_name"`
}
type ParentMenu struct{
    Url string `json:"url"`
    Title string `json:"title"`
    Module string `json:"module_name"`
    Action string `json:"method_name"`
    Children []Menu `json:"children,omitempty"`
}
func (ur *UserRole) GetMenus(role_id int32) (map[int32]ParentMenu,error) {
    menus := make(map[int32]ParentMenu)
    if role_id==0 {
        return menus,errors.New("Role ID Empty")
    }
    qb, _ := orm.NewQueryBuilder("mysql")
    var (
        ups []*SystemModule
    )
    table_module := new(SystemModule).TableName();
    table_access := new(UserPermission).TableName();
    qb.Select(table_module+".*").From( table_module ).
        LeftJoin( table_access ).On(table_module+".module_id = "+table_access+".module_id").
        Where("( "+ table_access +".role_id = ? OR parent_id=0 ) AND "+table_module+".show=1").
        OrderBy("type,sort")

    // 导出 SQL 语句
    sql := qb.String()

    // 执行 SQL 语句
    o := orm.NewOrm()
    _,err := o.Raw(sql, role_id).QueryRows(&ups)
    if err == nil {
        for _, up := range ups {
            if up.ParentId==0 {
                menus[up.ModuleId] = ParentMenu{"/"+up.ModuleName+"/"+up.MethodName,up.Title,up.ModuleName,up.MethodName,[]Menu{}}
            }else{
                var temp_data = ParentMenu{}
                if v,ok := menus[up.ParentId]; ok  {
                    temp_data = v
                }else{
                    continue
                }
                temp_data.Children = append( temp_data.Children,Menu{"/"+up.ModuleName+"/"+up.MethodName,up.Title,up.ModuleName,up.MethodName} )
                menus[up.ParentId] = temp_data
            }
        }
    }
    return menus,err
}

func (role *UserRole) Save() (bool,error){
    old_role := UserRole{RoleId:role.RoleId}
    o := orm.NewOrm()
    if role.RoleId!=0 && o.Read(&old_role) == nil {
        if _, err := o.Update(role); err != nil {
            return false,err
        }
    }else{
        _, err := o.Insert(role)
        if err != nil {
            return false,err
        }
    }
    return true,nil
}

func (role *UserRole) GetRoleList() (map[int32]UserRole,error){
    om := orm.NewOrm()
    roles := map[int32]UserRole{}
    tempData := []UserRole{}
    _,err := om.QueryTable(role).All( &tempData )
    if err == nil {
        for _, item := range tempData {
            item.Count,_ = om.QueryTable(new(SystemUser)).Filter("role_id",item.RoleId).Count()
            roles[item.RoleId] = item
        }
    }
    return roles,err
}

func (userp *UserPermission)GetAccessList(role_id int32)map[int32]int{
    om := orm.NewOrm()
    access := map[int32]int{}
    tempData := []UserPermission{}
    _,err := om.QueryTable(userp).Filter("RoleId",role_id).All( &tempData )
    if err == nil {
        for _, item := range tempData {
            access[item.ModuleId] = 1
        }
    }
    return access
}

func (role *UserRole) GetRoleName(role_id int32) string{
    var name = ""
    err := orm.NewOrm().QueryTable(role).Filter("RoleId",role_id).One(role)
    if err == nil {
        name = role.RoleName
    }
    return name
}

func (userp *UserPermission) SetPermission(role_id int32,permissios []UserPermission) bool  {
    if role_id==0 || permissios ==nil{
        return false
    }
    o := orm.NewOrm()
    if _, err := o.QueryTable(userp).Filter("RoleId", role_id).Delete(); err != nil {
        return false
    }
    if _, err := o.InsertMulti(100, permissios);err != nil {
        return false
    }
    return true
}
