// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package initializer

import "github.com/xfali/webproj/internal/pkg/value"

type Initializer func(target string, v *value.Value) error

var gInitializers = []Initializer{
    ProjectInitializer,
    MakeFileInitializer,
    AppInitializer,
    PackageInitializer,
    InternalInitializer,
}

func Initialize(target string, v *value.Value) error {
    for i := range gInitializers {
        err := gInitializers[i](target, v)
        if err != nil {
            return err
        }
    }
    return nil
}
