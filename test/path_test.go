// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
    "github.com/xfali/webproj/internal/pkg/initializer"
    "path/filepath"
    "testing"
)

func TestPath(t *testing.T) {
    path := " /x/x/x/ "
    t.Log(filepath.Clean(path))
    t.Log(initializer.FormatPath(path))

    path = " //x//x//x// "
    t.Log(filepath.Clean(path))
    t.Log(initializer.FormatPath(path))

    path = ` \x\x\x\ `
    t.Log(filepath.Clean(path))
    t.Log(initializer.FormatPath(path))

    path = ` \\x\\x\\x\\ `
    t.Log(filepath.Clean(path))
    t.Log(initializer.FormatPath(path))
}
