package libs

import "strings"

func UcFirst(str string) string {
    if str == ""{
        return ""
    }
    bytes := []byte(str)
    bytes[0] = []byte( strings.ToUpper(string(bytes[0])) )[0]
    return string(bytes)
}
