// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package tpltool

import (
    xio "github.com/xfali/goutils/io"
    "html/template"
    "io"
    "os"
    "path/filepath"
    "strings"
)

const (
    TempFileSuffix    = ".xtpl"
    TempFileSuffixLen = 5

    PublicApp   = "_your_public_app_"
    PublicPkg   = "_your_public_pkg_"
    PrivateApp = "_your_private_app_"
    PrivatePkg = "_your_private_pkg_"
)

func WalkFile(src, target string, handler func(src, target string) error) error {
    return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if info == nil {
            return err
        }
        if !info.IsDir() {
            destPath := strings.Replace(path, src, target, 1)
            if err := handler(path, destPath); err != nil {
                return err
            }
        }
        return nil
    })
}

func CopyTemplate(srcPath, destPath string, value interface{}) error {
    tpl, err := template.ParseFiles(srcPath)
    if err != nil {
        return err
    }

    destPath = FormatTemplateName(destPath)
    dest, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer dest.Close()
    return tpl.Execute(dest, value)
}

func CopyFile(srcPath, destPath string) error {
    src, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer src.Close()

    dest, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer dest.Close()

    _, errC := io.Copy(dest, src)
    return errC
}

func FormatTemplateName(filePath string) string {
    return filePath[:len(filePath)-TempFileSuffixLen]
}

func IsTemplateFile(filePath string) bool {
    pathLen := len(filePath)
    if pathLen <= TempFileSuffixLen {
        return false
    }

    if filePath[pathLen-TempFileSuffixLen:] == TempFileSuffix {
        return true
    } else {
        return false
    }
}

func TemplateDir(filePath string) string {
    dirName := filepath.Base(filePath)
    switch dirName {
    case PublicApp:
    case PublicPkg:
    case PrivateApp:
    case PrivatePkg:
        return dirName
    }
    return ""
}

func SafeMkdir(target string) error {
    if !xio.IsPathExists(target) {
        return xio.Mkdir(target)
    }
    return nil
}
