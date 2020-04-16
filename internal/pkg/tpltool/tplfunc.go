// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package tpltool

import "strings"

var tplFuncs = map[string]interface{}{
    "upperFirst": UpperFirst,
    "lowerFirst": LowerFirst,
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
