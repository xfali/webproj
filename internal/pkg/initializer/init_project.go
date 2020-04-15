// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package initializer

import (
    "github.com/xfali/goutils/io"
    "github.com/xfali/goutils/log"
    "github.com/xfali/webproj/internal/pkg/tpltool"
    "github.com/xfali/webproj/internal/pkg/value"
    "os"
    "path/filepath"
    "strings"
)

type ProjectInitializer struct{}

func NewProjectInitializer() *ProjectInitializer {
    return &ProjectInitializer{}
}

func (initer *ProjectInitializer) PreInit(ctx *value.Context) error {
    return nil
}

func (initer *ProjectInitializer) Init(src, target string, ctx *value.Context) error {
    if io.IsPathExists(target) {
        log.Fatal("Target Exists!")
    }
    err := tpltool.SafeMkdir(target)
    if err != nil {
        return err
    }

    return filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
        if info == nil {
            return err
        }
        fileName := filepath.Base(srcPath)
        srcPath = strings.Replace(srcPath, "\\", "/", -1)
        destPath := strings.Replace(srcPath, src, target, -1)
        if !info.IsDir() {
            if ctx.IgnoreMap[fileName] {
                return nil
            }
            if tpltool.IsTemplateFile(srcPath) {
                return tpltool.CopyTemplate(srcPath, destPath, ctx)
            } else {
                return tpltool.CopyFile(srcPath, destPath)
            }
        } else {
            if ctx.IgnoreMap[fileName] {
                return filepath.SkipDir
            }
            err := tpltool.SafeMkdir(destPath)
            if err != nil {
                return err
            }
        }
        return nil
    })
}

func (initer *ProjectInitializer) PostInit(src, target string, ctx *value.Context) error {
    return tpltool.Walk(target, func(srcPath *string, info os.FileInfo, err error) error {
        if info == nil {
            return err
        }

        originPath := *srcPath
        if info.IsDir() {
            dir, ok := ctx.DirMapper(originPath)
            if ok {
                if dir != "" {
                    *srcPath = filepath.Join(filepath.Dir(originPath), dir)
                    return os.Rename(originPath, *srcPath)
                } else {
                    *srcPath = ""
                    err := os.RemoveAll(originPath)
                    if err != nil {
                        return err
                    }
                    return filepath.SkipDir
                }
            }
        }
        return nil
    })
}
