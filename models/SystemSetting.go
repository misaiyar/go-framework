package models

import (
    "github.com/astaxie/beego/orm"
    "encoding/json"
)

func (config *SystemConfig) GetList(key string)([]SystemConfig,map[string]DicValue){
    om := orm.NewOrm()
    dicValueList := map[string]DicValue{}
    settings := []SystemConfig{}
    _,err := om.QueryTable(config).All( &settings )
    if err == nil {
        for index, item := range settings {
            if key!="" && item.Key == key || key=="" && index==0 {
                config.Key = item.Key
                config.Description = item.Description
                if item.Value !="" {
                    json.Unmarshal([]byte(item.Value), &dicValueList)
                }
            }
            settings[index].Value = ""
        }
    }
    return settings,dicValueList
}

func (config *SystemConfig) Save(add_dic_vals []DicValue,del_dic_vals []DicValue)(bool,error){
    old_config := SystemConfig{Key:config.Key}
    o := orm.NewOrm()
    dicValueList := make(map[string]DicValue)
    if o.Read(&old_config) == nil {
        if old_config.Value !="" {
            json.Unmarshal([]byte(old_config.Value), &dicValueList)
        }
        for _,dic_val := range add_dic_vals{
            dicValueList[dic_val.Key] = dic_val
        }
        for _,dic_val := range del_dic_vals{
            delete(dicValueList,dic_val.Key)
        }
        json,_ := json.Marshal(dicValueList)
        config.Value = string(json)
        if _, err := o.Update(config,"Value"); err != nil {
            return false,err
        }
    }else{
        _, err := o.Insert(config)
        if err != nil {
            return false,err
        }
    }
    return true,nil
}
