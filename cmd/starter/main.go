/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

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
    target := flag.String("o", ".", "output dir")
    flag.Parse()

    v := value.Read(*conf)
    if v == nil {
        os.Exit(-1)
    }

    err := initializer.Initialize(*target, v)
    if err != nil {
        log.Fatal(err)
    }
}
