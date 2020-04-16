// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
    "github.com/xfali/webproj/internal/pkg/tpltool"
    "testing"
)

func TestFunc(t *testing.T) {
    str := "abCdEfG"
    t.Log(tpltool.Camel2snake(str))
    t.Log(tpltool.Snake2camel(str))

    str = "AbCdEfG"
    t.Log(tpltool.Camel2snake(str))
    t.Log(tpltool.Snake2camel(str))

    str = "ab_cdE_fG"
    t.Log(tpltool.Camel2snake(str))
    t.Log(tpltool.Snake2camel(str))

    str = "Ab_CdE_fG"
    t.Log(tpltool.Camel2snake(str))
    t.Log(tpltool.Snake2camel(str))
}
