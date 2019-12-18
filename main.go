/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

package main

import (
    "github.com/gin-gonic/gin"
    "github.com/xfali/go-web-starter/web"
    "github.com/xfali/goutils/log"
    "github.com/xfali/restclient"
    "time"
)

type TestModel struct {
    T string
}

func main() {
    c := web.Startup(func(engine *gin.Engine) {
        engine.GET("/test", func(context *gin.Context) {
            context.JSON(200, TestModel{T: "test"})
        })
    })
    time.Sleep(time.Second)
    x := TestModel{}
    client := restclient.New()
    _, err := client.Get(&x, "http://localhost:8080/test")
    log.Info("err : %v ret : = %v", err, x)

    web.HandlerSignal(c)
}
