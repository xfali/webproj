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
    Project     string `json:"project"`
    Module      string `json:"module"`
    PublicApp   string `json:"publicApp"`
    PublicPkg   string `json:"publicPkg"`
    PrivateApp  string `json:"privateApp"`
    PrivatePkg  string `json:"privatePkg"`
    FileComment string `json:"fileComment"`
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
