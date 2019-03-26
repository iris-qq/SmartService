package config

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"path/filepath"
	"time"
)


func logWriter() io.Writer{
	BasePath,err:=filepath.Abs(filepath.Dir(os.Args[0]))
	if err !=nil{
		panic(err)
	}
	logPath:= filepath.Join(BasePath,"log","info.log")
	file,err:=os.Open(logPath)
	if err !=nil{
		fmt.Println(err)
	}
	return file
}

func InitConfig(e *echo.Echo){
	// 开启debug模式
	e.Debug = true
	// 设置log路径和级别
	logW:=logWriter()
	e.Logger.SetOutput(logW)
	e.Logger.SetLevel(log.DEBUG)
	// 隐藏echo横幅
	e.HideBanner=false
	// 设置超时时间为60秒
	e.Server.ReadTimeout=60*time.Second
	e.Static("/static","static")
}