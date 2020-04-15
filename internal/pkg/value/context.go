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

type Context struct {
    Value Value
}

func NewContext(valueFile string) *Context {
    ret := Context{
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
