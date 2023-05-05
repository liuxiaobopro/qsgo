package main

import (
	"embed"
	"fmt"
	"os"

	fmtp "github.com/liuxiaobopro/qsgo/log/fmt"
)

//go:embed tpl/*
var webTpls embed.FS

func initTpl() {
	// 创建.qsgo目录
	if _, err := os.Stat(qsgoPath); os.IsNotExist(err) {
		err = os.Mkdir(qsgoPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// 创建tpl目录, 覆盖生成
	if _, err := os.Stat(webTplPath); os.IsNotExist(err) {
		fmtp.Printf("%s不存在, 正在创建... \n", webTplPath)
		err = os.Mkdir(webTplPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmtp.Printf("%s创建成功 \n", webTplPath)
	} else {
		fmtp.Printf("%s存在, 正在删除... \n", webTplPath)
		err = os.RemoveAll(webTplPath)
		if err != nil {
			panic(err)
		}
		fmtp.Printf("%s删除成功, 正在创建... \n", webTplPath)
		err = os.Mkdir(webTplPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmtp.Printf("%s创建成功 \n", webTplPath)
	}

	// 获取projectWebTplPath下的所有文件
	files, err := webTpls.ReadDir(projectWebTplPath)
	if err != nil {
		panic(err)
	}

	// 将文件复制到webTplPath下
	for _, file := range files {
		// 读取文件
		data, err := webTpls.ReadFile(fmt.Sprintf("%s/%s", projectWebTplPath, file.Name()))
		if err != nil {
			panic(err)
		}

		// 写入文件
		if err := os.WriteFile(fmt.Sprintf("%s/%s", webTplPath, file.Name()), data, os.ModePerm); err != nil {
			panic(err)
		}
	}

}
