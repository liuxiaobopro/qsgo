package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/liuxiaobopro/qsgo/web"

	stringx "github.com/liuxiaobopro/gobox/string"
)

var (
	debug = false

	userHomePath string // 用户家目录
	qsgoPath     string
	webTplPath   string

	projectWebTplPath = "web/tpl"
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
	initWebTpl()
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
		return
	}

	for _, v := range args {
		if stringx.Has(v, byte('=')) {
			s := strings.Split(v, "=")
			switch s[0] {
			case "web":
				web.Start(s[1])
			}
		} else {
			if v == "h" || v == "help" {
				help()
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

用法：
  qsgo [参数]

参数：
  h, help  显示帮助信息[qsgo OR qsgo h OR qsgo help]
  web      生成web项目[qsgo web=项目名]
	`)
}
