package main

import (
	"fmt"
	"os"

	stringx "github.com/liuxiaobopro/gobox/string"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
		return
	}

	for _, v := range args {
		if stringx.Has(v, byte('=')) {
			if v == "web" {
				fmt.Println("web")
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
  qsgo [参数] 例如：qsgo h, qsgo web=Greet

参数：
  h, help  显示帮助信息
  web      生成web项目
	`)
}
