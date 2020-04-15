// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package main

import (
    "flag"
    "github.com/xfali/webproj/internal/pkg/initializer"
    "github.com/xfali/webproj/internal/pkg/value"
    "log"
    "os"
)

func main() {
    conf := flag.String("f", "value.json", "config file")
    src := flag.String("s", "./assets/project-layout", "source template dir")
    target := flag.String("o", ".", "output dir")
    flag.Parse()

    ctx := value.NewContext(*conf)
    if ctx == nil {
        os.Exit(-1)
    }


    err := initializer.Initialize(*src, *target, ctx)
    if err != nil {
        log.Fatal(err)
    }
}
