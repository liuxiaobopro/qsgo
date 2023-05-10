package web

import (
	"fmt"
	"os"
	"text/template"

	"github.com/liuxiaobopro/qsgo/global"
	fmtp "github.com/liuxiaobopro/qsgo/log/fmt"

	filex "github.com/liuxiaobopro/gobox/file"
	otherx "github.com/liuxiaobopro/gobox/other"
	stringx "github.com/liuxiaobopro/gobox/string"
)

var (
	webTplPath string
)

type ControllerLogic struct {
	Package   string
	Project   string
	Handle    string
	CL        string
	LogicPath string
	Logic     string
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
		fmtp.Println("获取项目名时出错：", err)
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

	fmtp.Println("controller目录路径:", controllerPath)
	fmtp.Println("logic目录路径:", logicPath)
	fmtp.Println("controller文件路径:", controllerFilePath)
	fmtp.Println("logic文件路径:", logicFilePath)
	fmtp.Printf("webTplPath: %s\n", webTplPath)

	// 判断controllerPath是否存在
	{
		if _, err := os.Stat(controllerFilePath); err == nil {
			fmtp.Printf("controller文件已存在: %s\n", controllerPath)
			return
		} else {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(controllerPath, os.ModePerm); err != nil {
					fmtp.Printf("创建controller目录失败: %s\n", err.Error())
					return
				}
				var file *os.File
				if file, err = os.Create(controllerFilePath); err != nil {
					fmtp.Printf("创建controller文件失败: %s\n", err.Error())
					return
				}

				// 解析模板文件
				data := ControllerLogic{
					Package:   getPackage(name, "controller"),
					Project:   global.ProjectName,
					Handle:    getHandle(name),
					CL:        getCL(name),
					Logic:     getLogic(name),
					LogicPath: getLogicPath(name),
				}
				tpl, err := template.ParseFiles(webTplPath + "/web_controller.tpl")
				if err != nil {
					fmtp.Printf("解析模板文件失败: %s\n", err.Error())
					return
				}
				// 应用模板，将结果写入新文件
				err = tpl.Execute(file, data)
				if err != nil {
					fmtp.Printf("应用模板失败: %s\n", err.Error())
					return
				}
			} else {
				fmtp.Printf("controller文件打开失败: %s\n", err.Error())
				return
			}
		}
	}

	// 判断logicPath是否存在
	{
		if _, err := os.Stat(logicFilePath); err == nil {
			fmtp.Printf("logic文件已存在: %s\n", logicPath)
			return
		} else {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(logicPath, os.ModePerm); err != nil {
					fmtp.Printf("创建logic目录失败: %s\n", err.Error())
					return
				}
				var file *os.File
				if file, err = os.Create(logicFilePath); err != nil {
					fmtp.Printf("创建logic文件失败: %s\n", err.Error())
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
					fmtp.Printf("解析模板文件失败: %s\n", err.Error())
					return
				}
				// 应用模板，将结果写入新文件
				err = tpl.Execute(file, data)
				if err != nil {
					fmtp.Printf("应用模板失败: %s\n", err.Error())
					return
				}
			} else {
				fmtp.Printf("logic文件打开失败: %s\n", err.Error())
				return
			}
		}
	}

	// 判断reply和req字符串是否存在
	{
		var (
			reqStruct                  string
			replyStruct                string
			reqFilePath, replyFilePath = targetPath + "define/types/req/", targetPath + "define/types/reply/"
		)
		if stringx.Has(name, byte('/')) {
			reqStruct = stringx.ReplaceCharAfterSpecifiedCharUp(name, "/") + "IndexReq"
			replyStruct = stringx.ReplaceCharAfterSpecifiedCharUp(name, "/") + "IndexReply"
		} else {
			n := stringx.FirstUp(name)
			reqStruct = n + "IndexReq"
			replyStruct = n + "IndexReply"
		}
		isHasStr1 := fmt.Sprintf("type %s struct", reqStruct)
		isHasStr2 := fmt.Sprintf("type %s struct", replyStruct)

		reqStruct = fmt.Sprintf("\n\ntype %s struct {}", reqStruct)
		replyStruct = fmt.Sprintf("\n\ntype %s struct {}", replyStruct)

		// 判断struct是否存在
		if has, err := filex.Has(reqFilePath+"req.go", isHasStr1); err != nil {
			fmtp.Printf("判断req.go文件是否存在失败: %s\n", err.Error())
			return
		} else {
			if !has {
				fmtp.Printf("req struct 不存在: %s\n", isHasStr1)
				// 追加内容
				if err := filex.Append(reqFilePath+"req.go", reqStruct); err != nil {
					fmtp.Printf("追加内容失败: %s\n", err.Error())
					return
				}
			} else {
				fmtp.Println("req struct 存在")
			}
		}

		// 判断struct是否存在
		if has, err := filex.Has(replyFilePath+"reply.go", isHasStr2); err != nil {
			fmtp.Printf("判断reply.go文件是否存在失败: %s\n", err.Error())
			return
		} else {
			if !has {
				fmtp.Printf("reply struct 不存在: %s\n", isHasStr2)
				// 追加内容
				if err := filex.Append(replyFilePath+"reply.go", replyStruct); err != nil {
					fmtp.Printf("追加内容失败: %s\n", err.Error())
					return
				}
			} else {
				fmtp.Println("reply struct 存在")
			}
		}
	}

	fmt.Println("Done!")
}
