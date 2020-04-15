// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package initializer

import "github.com/xfali/webproj/internal/pkg/value"

type Initializer func(src, target string, ctx *value.Context) error

var gInitializers = []Initializer{
    ProjectInitializer,
    MakeFileInitializer,
    AppInitializer,
    PackageInitializer,
    InternalInitializer,
}

func Initialize(src, target string, ctx *value.Context) error {
    for i := range gInitializers {
        err := gInitializers[i](src, target, ctx)
        if err != nil {
            return err
        }
    }
    return nil
}
