package libs

import (
    "time"
    "github.com/astaxie/beego/httplib"
    "github.com/astaxie/beego/logs"
)

func Curl_get(url string, params map[string]string) (string,error) {
    req := httplib.Get(url).SetTimeout(2*time.Second, 3*time.Second)
    for k, v := range params {
        req.Param(k, v)
    }
    str, err := req.String()
    if err != nil {
        logs.Error(err.Error())
        return "",err
    }
    return str,nil
}

func Curl_post(url string, params map[string]string) (string,error) {
    req := httplib.Post(url).SetTimeout(2*time.Second, 3*time.Second)
    for k, v := range params {
        req.Param(k, v)
    }
    str, err := req.String()
    if err != nil {
        logs.Error(err.Error())
        return "",err
    }
    return str,nil
}
