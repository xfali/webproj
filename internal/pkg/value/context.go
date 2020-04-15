// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package value

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "strings"
)

type Context struct {
    Value Value

    IgnoreMap map[string]bool
}

func NewContext(valueFile string) *Context {
    ret := Context{
        IgnoreMap: map[string]bool{
            ".git":  true,
            ".idea": true,
        },
    }
    err := ret.ReadValue(valueFile)
    if err != nil {
        log.Fatal(err)
    }
    return &ret
}

func (c *Context) ReadValue(valueFile string) error {
    b, err := ioutil.ReadFile(valueFile)
    if err != nil {
        return err
    }
    return json.Unmarshal(b, &c.Value)
}

func (c *Context) AddIgnore(ig string) {
    list := strings.Split(ig, "|")
    for _, v := range list {
        v = strings.TrimSpace(v)
        c.IgnoreMap[v] = true
    }
}
