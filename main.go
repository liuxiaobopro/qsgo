package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/liuxiaobopro/qsgo/global"
	fmtp "github.com/liuxiaobopro/qsgo/log/fmt"
	"github.com/liuxiaobopro/qsgo/service/web"

	arrayx "github.com/liuxiaobopro/gobox/array"
	stringx "github.com/liuxiaobopro/gobox/string"
)

var (
	version string
	debug   = false

	userHomePath string // 用户家目录
	qsgoPath     string
	webTplPath   string

	projectWebTplPath = "tpl"

	//go:embed VERSION
	versionFile embed.FS
)

func init() {
	//#region 初始化
	//#region 通用
	global.Debug = debug
	global.Version = version
	//#endregion
	//#region 用户相关
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fmtp.Println("home:", u)
	global.UserHomePath, userHomePath = u, u
	global.QsgoPath, qsgoPath = userHomePath+"/.qsgo", userHomePath+"/.qsgo"
	global.WebTplPath, webTplPath = qsgoPath+"/tpl", qsgoPath+"/tpl"
	//#endregion

	//#region 项目相关
	// 获取当前目录
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmtp.Println("pwd:", pwd)
	global.ProjectPath = pwd
	//#endregion
	//#endregion

	//#region 获取当前目录下version文件的内容
	versionByte, err := versionFile.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}
	version = string(versionByte)
	//#endregion

	initTpl()
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
		return
	}

	if len(args) == 2 {
		if args[0] == "debug" {
			debug = true
			global.Debug = debug
		}
	}

	var v string
	if arrayx.IsIn(args, "debug") {
		debug = true
		global.Debug = debug

		if len(args) <= 1 {
			fmt.Println("debug不能单独使用")
			return
		}

		v = args[1]
	} else {
		v = args[0]
	}

	fmtp.Println("args:", strings.Join(args, " "))

	if stringx.Has(v, byte(':')) {
		s := strings.Split(v, ":")
		switch s[0] {
		case "web":
			web.Start(s[1])
		}
	} else {
		switch v {
		case "h":
		case "help":
			help()
		default:
			fmt.Println("参数错误")
		}
	}
}

func help() {
	fmt.Println(`
	 _______  _______  _______  _______
	(  ___  )(  ____ \(  ____ \(  ___  )
	| (   ) || (    \/| (    \/| (   ) |
	| |   | || (_____ | |      | |   | |
	| |   | |(_____  )| | ____ | |   | |
	| | /\| |      ) || | \_  )| |   | |
	| (_\ \ |/\____) || (___) || (___) |
	(____\/_)\_______)(_______)(_______)

	 这是一个辅助完成Go项目的工具

	 Version: ` + version + `

用法：
  qsgo [参数:功能=参数值]

参数：
---------------------------------------------------
  h, help  显示帮助信息
    qsgo
    qsgo h
    qsgo help
---------------------------------------------------
  web      生成web项目

    qsgo web:name=项目名 [生成项目]
	  - **必须安装git**
	  - 最好找个空目录执行
	  - 第一次执行可能会失败, 请再次执行

    qsgo web:api=接口名 [生成接口]
	  - **必须在项目目录下执行**
	`)
}
