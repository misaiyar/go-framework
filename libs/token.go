package libs

import (
    "crypto/md5"
    "encoding/hex"
    "sort"
    "strings"
    "fmt"
)

var AuthApp = map[string]string{"app_key":"app_secret"}

func Look_array_token(param map[string]string) string {
    var keys []string
    for k := range param {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    token := ""
    for _, k := range keys {
        token += k + param[k] + k + param[k]
    }
    return token
}
func GenToken(secretkey string, param map[string]string) string {
    token := secretkey
    token += Look_array_token(param)
    token += secretkey
    fmt.Println(token)
    h := md5.New()
    h.Write([]byte(token)) // 需要加密的字符串
    token = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
    fmt.Println(token)
    return token
}
