package models

import (
    "github.com/astaxie/beego/orm"
    "project/libs"
)

func (module *SystemModule) GetModuleList() (map[int32]SystemModule,error){
    om := orm.NewOrm()
    modules := map[int32]SystemModule{}
    tempData := []SystemModule{}
    _,err := om.QueryTable(module).OrderBy("parent_id","sort").All( &tempData )
    if err == nil {
        for _, item := range tempData {
            if item.ParentId==0 {
                item.Children = make(map[int32]SystemModule)
                modules[item.ModuleId] = item
            }else{
                modules[item.ParentId].Children[item.ModuleId] = item

            }
        }
    }
    return modules,err
}

func (module *SystemModule) Save() (bool,error){
    old_module := SystemModule{ModuleId:module.ModuleId}
    if module.Type == 1 {
        module.MethodName = ""
    }
    o := orm.NewOrm()
    if module.ModuleId!=0 && o.Read(&old_module) == nil {
        if _, err := o.Update(module); err != nil {
            return false,err
        }
    }else{
        _, err := o.Insert(module)
        if err != nil {
            return false,err
        }
    }
    return true,nil
}

func (module *SystemModule)GetAccessModuleList(role_id int32)map[string]map[string]bool{
    qb, _ := orm.NewQueryBuilder("mysql")
    access := map[string]map[string]bool{}
    tempData := []SystemModule{}
    table_module := module.TableName();
    table_access := new(UserPermission).TableName();
    qb.Select(table_module+".*").From( table_module ).
        LeftJoin( table_access ).On(table_module+".module_id = "+table_access+".module_id").
        Where("( "+ table_access +".role_id = ? )")

    // 导出 SQL 语句
    sql := qb.String()

    // 执行 SQL 语句
    o := orm.NewOrm()
    _,err := o.Raw(sql, role_id).QueryRows(&tempData)
    if err == nil {
        for _, item := range tempData {
            item.ModuleName = libs.UcFirst(item.ModuleName)
            item.MethodName = libs.UcFirst(item.MethodName)
            if _ ,ok:=access[item.ModuleName];ok {
                access[item.ModuleName][item.MethodName] = true
            }else{
                access[item.ModuleName] = map[string]bool{item.MethodName:true}
            }
        }
    }
    return access
}
