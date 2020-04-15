// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package value

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"
)

const (
    PublicApp  = "_your_public_app_"
    PublicPkg  = "_your_public_lib_"
    PrivateApp = "_your_private_app_"
    PrivatePkg = "_your_private_lib_"
)

type Context struct {
    Value Value

    IgnoreMap map[string]bool
    DirMap    map[string]string
}

func NewContext(valueFile string) *Context {
    ret := Context{
        IgnoreMap: map[string]bool{
            ".git":  true,
            ".idea": true,
        },
        DirMap: map[string]string{},
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
    err = json.Unmarshal(b, &c.Value)
    if err != nil {
        return err
    }

    c.DirMap[PublicApp] = c.Value.PublicApp
    c.DirMap[PublicPkg] = c.Value.PublicPkg
    c.DirMap[PrivateApp] = c.Value.PrivateApp
    c.DirMap[PrivatePkg] = c.Value.PrivateApp
    return nil
}

func (c *Context) AddIgnore(ig string) {
    list := strings.Split(ig, "|")
    for _, v := range list {
        v = strings.TrimSpace(v)
        c.IgnoreMap[v] = true
    }
}

func (c *Context) DirMapper(dir string) (string, bool) {
    v, ok := c.DirMap[filepath.Base(dir)]
    return v, ok
}
