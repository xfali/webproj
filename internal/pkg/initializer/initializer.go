// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package initializer

import (
    "github.com/xfali/webproj/internal/pkg/value"
    "strings"
    "unicode"
)

type Initializer interface {
    PreInit(ctx *value.Context) error
    Init(src, target string, ctx *value.Context) error
    PostInit(src, target string, ctx *value.Context) error
}

var gInitializers = []Initializer{
    NewProjectInitializer(),
}

func Initialize(src, target string, ctx *value.Context) (err error) {
    src = FormatPath(src)
    target = FormatPath(target)

    initializer := NewProjectInitializer()
    err = initializer.PreInit(ctx)
    if err != nil {
        return
    }
    err = initializer.Init(src, target, ctx)
    if err != nil {
        return
    }
    err = initializer.PostInit(src, target, ctx)
    return
}

func FormatPath(src string) string {
    return strings.TrimRightFunc(strings.TrimLeftFunc(src, unicode.IsSpace), func(r rune) bool {
        return unicode.IsSpace(r) || r == '/' || r == '\\'
    })
}
