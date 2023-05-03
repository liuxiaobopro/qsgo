package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/liuxiaobopro/qsgo/service/web"

	stringx "github.com/liuxiaobopro/gobox/string"
)

var (
	version string
	debug   = true

	userHomePath string // 用户家目录
	qsgoPath     string
	webTplPath   string

	projectWebTplPath = "tpl"

	//go:embed VERSION
	versionFile embed.FS
)

func init() {
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if debug {
		fmt.Println("home:", u)
	}
	userHomePath = u
	qsgoPath = userHomePath + "/.qsgo"
	webTplPath = qsgoPath + "/tpl"

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

	for _, v := range args {
		if stringx.Has(v, byte(':')) {
			s := strings.Split(v, ":")
			switch s[0] {
			case "web":
				web.Start(s[1])
			}
		} else {
			if v == "h" || v == "help" {
				help()
			} else {
				fmt.Println("参数错误")
			}
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
  	生成项目  qsgo web:name=项目名
  	生成接口  qsgo web:api=接口名
	`)
}
