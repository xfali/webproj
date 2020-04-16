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
    "text/template"
)

const (
    PublicApp  = "_your_public_app_"
    PublicPkg  = "_your_public_lib_"
    PrivateApp = "_your_private_app_"
    PrivatePkg = "_your_private_lib_"

    DirTemplatePrefix    = "_xtpl_"
    DirTemplatePrefixLen = 6
)

type Context struct {
    Value map[string]interface{}

    IgnoreMap map[string]bool
    DirMap    map[string]string
}

func NewContext(valueFile string) *Context {
    ret := Context{
        Value: map[string]interface{}{},
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

    c.DirMap[PublicApp] = c.Value["PublicApp"].(string)
    c.DirMap[PublicPkg] = c.Value["PublicPkg"].(string)
    c.DirMap[PrivateApp] = c.Value["PrivateApp"].(string)
    c.DirMap[PrivatePkg] = c.Value["PrivateApp"].(string)
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
    //v, ok := c.DirMap[filepath.Base(dir)]
    dir = filepath.Base(dir)
    dirLen := len(dir)
    if dirLen <= DirTemplatePrefixLen {
        return "", false
    }

    if dir[:DirTemplatePrefixLen] == DirTemplatePrefix {
        tplDir := dir[DirTemplatePrefixLen:]
        vKeys := strings.Split(tplDir, "_")
        if len(vKeys) == 1 || vKeys[0] != "Value" {
            return "", true
        }
        templateStr := "{{"
        for _, v := range vKeys {
            templateStr += "." + v
        }
        templateStr += "}}"
        tpl, err := template.New("").Option("missingkey=zero").Parse(templateStr)
        if err != nil {
            return "", true
        }
        b := strings.Builder{}
        err = tpl.Execute(&b, c)
        if err != nil {
            return "", true
        }
        return b.String(), true
    }
    return "", false
}
