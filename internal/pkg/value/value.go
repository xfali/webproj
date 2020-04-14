// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package value

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

type Value struct {
    Project     string   `json:"project"`
    Module      string   `json:"module"`
    App         []string `json:"app"`
    Pkg         []string `json:"pkg"`
    InternalApp []string `json:"internalApp"`
    InternalPkg []string `json:"internalPkg"`
}

func Read(path string) *Value {
    b, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    v := Value{}
    err = json.Unmarshal(b, &v)
    if err != nil {
        log.Fatal(err)
    }

    return &v
}
