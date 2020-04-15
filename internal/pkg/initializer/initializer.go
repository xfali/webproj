// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package initializer

import (
    "github.com/xfali/webproj/internal/pkg/value"
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
