// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package tpltool

import (
    "github.com/Masterminds/sprig"
    xio "github.com/xfali/goutils/io"
    "html/template"
    "io"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

const (
    TempFileSuffix    = ".xtpl"
    TempFileSuffixLen = 5
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
    tpl := template.New(filepath.Base(srcPath)).Funcs(tplFuncs)
    tpl = tpl.Funcs(sprig.FuncMap())
    tpl, err := tpl.ParseFiles(srcPath)
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

func SafeMkdir(target string) error {
    if !xio.IsPathExists(target) {
        return xio.Mkdir(target)
    }
    return nil
}

func Walk(root string, walkFn WalkFunc) error {
    info, err := os.Lstat(root)
    if err != nil {
        err = walkFn(&root, nil, err)
    } else {
        err = walk(&root, info, walkFn)
    }
    if err == filepath.SkipDir {
        return nil
    }
    return err
}

var lstat = os.Lstat // for testing

type WalkFunc func(path *string, info os.FileInfo, err error) error

// walk recursively descends path, calling walkFn.
func walk(path *string, info os.FileInfo, walkFn WalkFunc) error {
    if path == nil || *path == "" {
        return nil
    }

    if !info.IsDir() {
        return walkFn(path, info, nil)
    }

    names, err := readDirNames(*path)
    err1 := walkFn(path, info, err)
    // If err != nil, walk can't walk into this directory.
    // err1 != nil means walkFn want walk to skip this directory or stop walking.
    // Therefore, if one of err and err1 isn't nil, walk will return.
    if err != nil || err1 != nil {
        // The caller's behavior is controlled by the return value, which is decided
        // by walkFn. walkFn may ignore err and return nil.
        // If walkFn returns SkipDir, it will be handled by the caller.
        // So walk should return whatever walkFn returns.
        return err1
    }

    for _, name := range names {
        filename := filepath.Join(*path, name)
        fileInfo, err := lstat(filename)
        if err != nil {
            if err := walkFn(&filename, fileInfo, err); err != nil && err != filepath.SkipDir {
                return err
            }
        } else {
            err = walk(&filename, fileInfo, walkFn)
            if err != nil {
                if !fileInfo.IsDir() || err != filepath.SkipDir {
                    return err
                }
            }
        }
    }
    return nil
}

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entries.
func readDirNames(dirname string) ([]string, error) {
    f, err := os.Open(dirname)
    if err != nil {
        return nil, err
    }
    names, err := f.Readdirnames(-1)
    f.Close()
    if err != nil {
        return nil, err
    }
    sort.Strings(names)
    return names, nil
}
