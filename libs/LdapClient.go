package libs

import (
    "gopkg.in/ldap.v2"
    "fmt"
    "errors"
)

type LdapInfo struct {
    UserName string
    Email string
    RealName string
}
var LdapHost = ""
var LdapPort = 0
var LdapDn = ""
func SearchLdap(username string,passwd string) (LdapInfo,error){
    l,err := ldap.Dial("tcp",fmt.Sprintf("%s:%d",LdapHost,LdapPort));
    if err !=nil {
        return LdapInfo{},err
    }
    defer l.Close()

    filter := fmt.Sprintf("(|(uid=%s)(mail=%s))",username,username)
    searchRequest := ldap.NewSearchRequest(LdapDn,ldap.ScopeWholeSubtree,
        ldap.NeverDerefAliases,0,0,false,filter,[]string{"dn","cn","uid","mail"},nil)
    sr,err := l.Search(searchRequest)
    if err !=nil {
        return LdapInfo{},err
    }
    if len(sr.Entries) != 1{
        return LdapInfo{},errors.New("User does not exist or too many entries returned")
    }
    userdn := sr.Entries[0].DN
    err = l.Bind(userdn,passwd)
    if err != nil {
        return LdapInfo{},err
    }
    return LdapInfo{sr.Entries[0].GetAttributeValue("uid"),
        sr.Entries[0].GetAttributeValue("mail"),sr.Entries[0].GetAttributeValue("cn")},nil
}
