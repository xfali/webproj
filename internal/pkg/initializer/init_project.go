// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package initializer

import (
    "github.com/xfali/webproj/internal/pkg/tpltool"
    "github.com/xfali/webproj/internal/pkg/value"
    "os"
    "path/filepath"
    "strings"
)

func ProjectInitializer(src, target string, ctx *value.Context) error {
    err := tpltool.SafeMkdir(target)
    if err != nil {
        return err
    }

    return filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
        if info == nil {
            return err
        }
        srcPath = strings.Replace(srcPath, "\\", "/", -1)
        destPath := strings.Replace(srcPath, src, target, -1)
        if !info.IsDir() {
            if tpltool.IsTemplateFile(srcPath) {
                return tpltool.CopyTemplate(srcPath, destPath, ctx)
            } else {
                return tpltool.CopyFile(srcPath, destPath)
            }
        } else {
            err := tpltool.SafeMkdir(destPath)
            if err != nil {
                return err
            }
        }
        return nil
    })
}
