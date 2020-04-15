# webproj

## 介绍

webproj是一个根据配置自动生成go project layout的工具。

项目的目录参考[project-layout](https://github.com/golang-standards/project-layout)

## 安装
```
go get github.com/xfali/webproj/cmd/starter
```

## 使用
```
./starter -f ${VALUE_FILE} -s ${TEMPLATE_PROJ_DIR} -o ${TARGET_DIR}
```

## 说明
### value
目前读取的value文件格式为json，可在文件中增加自定义的值

默认的[value文件](configs/value.json)：
```
{
  "Project": "PROJECT-NAME",
  "Module": "github.com/USER-ORG-NAME/REPO-NAME",
  "PublicApp": "YOUR-APP",
  "PublicPkg": "YOUR-PACKAGE",
  "PrivateApp": "",
  "PrivatePkg": "",
  "FileComment": "// Copyright (C) 2020, YOUR-ORG-NAME.\n// @author YOUR-NAME\n// @version YOUR-VERSION\n// Description: \n\n"
}
```
### 在项目模板中使用value
1. 识别的模板文件后缀名为.xtpl
2. 模板文件遵循go template规范
3. 使用value： 
```
{{.Value.Project}}
{{.Value.YOUR_VALUE}}
```

## 内置项目模板
位于assets/project-layout

[Default project template](assets/project-layout)

## 更多项目模板
开发中...

