// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
    "encoding/json"
    "os"
    "testing"
    "text/template"
)

var jsonStr = `
{
    "i" : 1,
    "s" : "test",
    "o" : {
        "f1" :  "name"
    }
}
`

var tplStr = `
value i : {{.i}}

value s : {{.s}}

value o : {{.o}}

value o.f1 : {{.o.f1}}
`

func TestValue(t *testing.T) {
    v := map[string]interface{}{}
    err := json.Unmarshal([]byte(jsonStr), &v)
    if err != nil {
        t.Fatal(err)
    }

    t.Log("i not match: ", v["i"])

    if v["s"].(string) != "test" {
        t.Fatal("s not match: ", v["s"])
    } else {
        t.Log("s = test")
    }

    if m, ok := v["o"].(map[string]interface{}); ok {
        if m["f1"].(string) != "name" {
            t.Fatal(`o["f1"] not match: `, m["f1"])
        } else {
            t.Log("o.f = name")
        }
    } else {
        t.Fatal("o not match: ", v["o"])
    }
}

func TestTemplate(t *testing.T) {
    v := map[string]interface{}{}
    err := json.Unmarshal([]byte(jsonStr), &v)
    if err != nil {
        t.Fatal(err)
    }

    tpl := template.New("test")
    tpl, err = tpl.Parse(tplStr)
    if err != nil {
        t.Fatal(err)
    }
    err = tpl.Execute(os.Stdout, v)
    if err != nil {
        t.Fatal(err)
    }
}
