package web

import (
	"fmt"
	"os"
	"sync"
	"text/template"

	"github.com/liuxiaobopro/qsgo/global"

	otherx "github.com/liuxiaobopro/gobox/other"
	stringx "github.com/liuxiaobopro/gobox/string"
)

var (
	webTplPath string
)

type ControllerLogic struct {
	Package string
	Project string
	Handle  string
	CL      string
}

func init() {
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	webTplPath = u + "/.qsgo/tpl"

	// 获取当前目录
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var err1 error
	global.ProjectName, err1 = otherx.GetProjectName(pwd)
	if err1 != nil {
		fmt.Println("获取项目名时出错：", err)
		return
	}
}

func api(name string) {
	fmt.Printf("开始创建api: %s\n", name)
	// 目标目录
	targetPath := "./"

	controllerPath := targetPath + "controller/" + name
	controllerPath = stringx.CutStartString(controllerPath, '/')
	controllerFilePath := targetPath + "controller/" + name + ".go"
	logicPath := targetPath + "logic/" + name
	logicPath = stringx.CutStartString(logicPath, '/')
	logicFilePath := targetPath + "logic/" + name + ".go"

	if global.Debug {
		fmt.Println("controller目录路径:", controllerPath)
		fmt.Println("logic目录路径:", logicPath)
		fmt.Println("controller文件路径:", controllerFilePath)
		fmt.Println("logic文件路径:", logicFilePath)
		fmt.Printf("webTplPath: %s\n", webTplPath)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// 判断controllerPath是否存在
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if _, err := os.Stat(controllerFilePath); err == nil {
			fmt.Printf("controller文件已存在: %s\n", controllerPath)
			return
		} else {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(controllerPath, os.ModePerm); err != nil {
					fmt.Printf("创建controller目录失败: %s\n", err.Error())
					return
				}
				var file *os.File
				if file, err = os.Create(controllerFilePath); err != nil {
					fmt.Printf("创建controller文件失败: %s\n", err.Error())
					return
				}

				// 解析模板文件
				data := ControllerLogic{
					Package: getPackage(name, "controller"),
					Project: global.ProjectName,
					Handle:  getHandle(name),
					CL:      getCL(name),
				}
				tpl, err := template.ParseFiles(webTplPath + "/web_controller.tpl")
				if err != nil {
					fmt.Printf("解析模板文件失败: %s\n", err.Error())
					return
				}
				// 应用模板，将结果写入新文件
				err = tpl.Execute(file, data)
				if err != nil {
					fmt.Printf("应用模板失败: %s\n", err.Error())
					return
				}
			} else {
				fmt.Printf("controller文件打开失败: %s\n", err.Error())
				return
			}
		}
	}(&wg)

	// 判断logicPath是否存在
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if _, err := os.Stat(logicFilePath); err == nil {
			fmt.Printf("logic文件已存在: %s\n", logicPath)
			return
		} else {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(logicPath, os.ModePerm); err != nil {
					fmt.Printf("创建logic目录失败: %s\n", err.Error())
					return
				}
				var file *os.File
				if file, err = os.Create(logicFilePath); err != nil {
					fmt.Printf("创建logic文件失败: %s\n", err.Error())
					return
				}

				// 解析模板文件
				data := ControllerLogic{
					Package: getPackage(name, "logic"),
					Project: global.ProjectName,
					Handle:  getHandle(name),
					CL:      getCL(name),
				}
				tpl, err := template.ParseFiles(webTplPath + "/web_logic.tpl")
				if err != nil {
					fmt.Printf("解析模板文件失败: %s\n", err.Error())
					return
				}
				// 应用模板，将结果写入新文件
				err = tpl.Execute(file, data)
				if err != nil {
					fmt.Printf("应用模板失败: %s\n", err.Error())
					return
				}
			} else {
				fmt.Printf("logic文件打开失败: %s\n", err.Error())
				return
			}
		}
	}(&wg)

	wg.Wait()
}

func getPackage(name, d string) string {
	if !stringx.Has(name, byte('/')) {
		return d
	} else {
		return stringx.CutEndString(name, '/')
	}
}

func getHandle(name string) string {
	return stringx.ReplaceCharAfterSpecifiedCharLow(name, "/")
}

func getCL(name string) string {
	return stringx.ReplaceCharAfterSpecifiedCharUp(name, "/")
}
