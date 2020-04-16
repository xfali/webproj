// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package tpltool

import "strings"

var tplFuncs = map[string]interface{}{
    "upper_first": UpperFirst,
    "lower_first": LowerFirst,
    "camel2snake": Camel2snake,
    "snake2camel": Snake2camel,
}

func Camel2snake(s string) string {
    data := make([]byte, 0, len(s)*2)
    j := false
    num := len(s)
    for i := 0; i < num; i++ {
        d := s[i]
        if i > 0 && d >= 'A' && d <= 'Z' && j {
            data = append(data, '_')
        }
        if d != '_' {
            j = true
        }
        data = append(data, d)
    }
    return strings.ToLower(string(data))
}

// camel string, xx_yy to XxYy
func Snake2camel(s string) string {
    data := make([]byte, 0, len(s))
    j := false
    k := false
    num := len(s) - 1
    for i := 0; i <= num; i++ {
        d := s[i]
        if k == false && d >= 'A' && d <= 'Z' {
            k = true
        }
        if d >= 'a' && d <= 'z' && (j || k == false) {
            d = d - 32
            j = false
            k = true
        }
        if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
            j = true
            continue
        }
        data = append(data, d)
    }
    return string(data)
}

func UpperFirst(src string) string {
    if src == "" {
        return src
    }
    return strings.ToUpper(src[:1]) + src[1:]
}

func LowerFirst(src string) string {
    if src == "" {
        return src
    }
    return strings.ToLower(src[:1]) + src[1:]
}
