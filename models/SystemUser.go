package models

import (
    "github.com/astaxie/beego/orm"
    "project/libs"
    "errors"
    "crypto/md5"
    "encoding/hex"
)

func (user *SystemUser) Insert(name string,passwd string,email string,realname string) (int64,error){
    om := orm.NewOrm()
    user.UserName = name
    h := md5.New()
    h.Write([]byte(passwd)) // 需要加密的字符串
    passwd_md5 := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
    user.Password = passwd_md5
    user.RealName = realname
    user.Email = email
    return om.Insert(user)
}

func (user *SystemUser) CheckLogin(username string,passwd string) (bool,error){
    ldapInfo,err := libs.SearchLdap( username,passwd )
    if err !=nil {//ldap不通过，使用表数据
        om := orm.NewOrm()
        conds := orm.NewCondition()
        cond := conds.Or("username",username).Or("email",username)
        err := om.QueryTable(user).SetCond(cond).One(user)
        if err !=nil {
            return false,errors.New("用户不存在")
        }
        h := md5.New()
        h.Write([]byte(passwd)) // 需要加密的字符串
        passwd_md5 := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
        if( user.Password != passwd_md5 ){
            return false,errors.New("密码错误")
        }
    } else {
        o := orm.NewOrm()
        user.UserName = ldapInfo.UserName
        user.RealName = ldapInfo.RealName
        user.Email = ldapInfo.Email
        user.Status = 1
        // 三个返回参数依次为：是否新创建的，对象 Id 值，错误
        if _, _, err := o.ReadOrCreate(user, "Email"); err != nil {
            return false,errors.New("创建用户失败")
        }
    }
    return true,nil
}

func (user *SystemUser) GetRoleUsers(role_id int32) map[string]string{
    var roleUsers = make(map[string]string)
    if role_id == 0 {
        return roleUsers
    }
    userList := []SystemUser{}
    om := orm.NewOrm()
    _,err := om.QueryTable(user).Filter("role_id",role_id).All(&userList);
    if err == nil {
        for _, item := range userList {
            roleUsers[item.UserName] = item.RealName
        }
    }
    return  roleUsers
}

func (user *SystemUser)GetPage(page int32,limit int32,baseUrl string,username string) ([]SystemUser,Pagination)  {
    om := orm.NewOrm();
    userList := []SystemUser{};
    query := om.QueryTable(user);
    if username != "" {
        query = query.Filter("Username", username)
    }
    count,err := query.Count();
    if err !=nil {
        count = 0;
    }
    query.Offset((page-1)*limit).Limit(limit).All(&userList,"UserId","UserName","Email","RealName","RoleId","Status","Created")
    return userList, Pagination{CurrentPage:page,TotalCounts:int32(count),BaseUrl:baseUrl}
}


func (user *SystemUser) Save() (bool,error){
    old_user := SystemUser{UserId:user.UserId}
    o := orm.NewOrm()
    if user.UserId!=0 && o.Read(&old_user) == nil {
        if _, err := o.Update(user,"Email","Realname","Status","RoleId"); err != nil {
            return false,err
        }
    }else{
        _, err := o.Insert(user)
        if err != nil {
            return false,err
        }
    }
    return true,nil
}
func (user *SystemUser) Update(field ...string) (bool,error){
    o := orm.NewOrm()
    if _, err := o.Update(user,field...); err != nil {
        return false,err
    }
    return true,nil
}
