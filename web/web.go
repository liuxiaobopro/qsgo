package web

import (
	"fmt"
	"os"
	"text/template"
)

var (
	tplPath string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	tplPath = fmt.Sprintf("%s/.qsgo/web/", home)
}

func Start(arg string) {
	// 定义要传递给模板的数据
	data := struct {
		AuthImport string
	}{
		AuthImport: "myauth",
	}

	// 解析模板文件
	tpl, err := template.ParseFiles(tplPath + "demo.tpl")
	if err != nil {
		panic(err)
	}

	// 创建新文件
	file, err := os.Create(fmt.Sprintf("temp/%s.go", arg))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 应用模板，将结果写入新文件
	err = tpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
