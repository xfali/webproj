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
    Project     string `json:"Project"`
    Module      string `json:"Module"`
    PublicApp   string `json:"PublicApp"`
    PublicPkg   string `json:"PublicPkg"`
    PrivateApp  string `json:"PrivateApp"`
    PrivatePkg  string `json:"PrivatePkg"`
    FileComment string `json:"FileComment"`
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
