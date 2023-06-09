package web

import (
	"fmt"

	fmtp "github.com/liuxiaobopro/qsgo/log/fmt"

	filex "github.com/liuxiaobopro/gobox/file"
	stringx "github.com/liuxiaobopro/gobox/string"
)

func router(name string) {
	fmt.Printf("开始创建router: %s\n", name) // demo/UserInfo
	if !stringx.Has(name, '/') {
		fmt.Printf("router格式错误,例如: demo/%s, 就是在demo.go添加%s方法", name, name)
		return
	}

	path := stringx.CutStartString(name, '/')
	path = path[:len(path)-1]
	fmtp.Println("path:", path)
	handleFunc := stringx.CutEndString(name, '/')
	handle := stringx.ReplaceCharAfterSpecifiedCharLow(path, "/")
	logicVar := stringx.ReplaceCharAfterSpecifiedCharUp(path, "/")
	var logic = "logic"
	fmtp.Println("handle:", handle)
	if stringx.Has(path, byte('/')) {
		logic = getLogic(path)
	}

	controllerFilePath := fmt.Sprintf("./controller/%s.go", path)
	logicFilePath := fmt.Sprintf("./logic/%s.go", path)
	reqFilePath := "./define/types/req/req.go"
	replyFilePath := "./define/types/reply/reply.go"

	router := &genRouter{
		Func:   handleFunc,
		Handle: handle,
		CL:     getCL(name),
		Logic:  logic,
		LogicVar:logicVar,
	}

	if has, err := router.checkRouterController(controllerFilePath); err != nil {
		fmt.Printf("检查controller失败: %v\n", err)
		return
	} else {
		if has {
			fmt.Printf("controller已存在: %s\n", controllerFilePath)
		} else {
			if err := filex.Append(controllerFilePath, router.genRouterController()); err != nil {
				fmt.Printf("创建controller失败: %v\n", err)
				return
			}
		}
	}

	if has, err := router.checkRouterLogic(logicFilePath); err != nil {
		fmt.Printf("检查logic失败: %v\n", err)
		return
	} else {
		if has {
			fmt.Printf("logic已存在: %s\n", logicFilePath)
		} else {
			if err := filex.Append(logicFilePath, router.genRouterLogic()); err != nil {
				fmt.Printf("创建logic失败: %v\n", err)
				return
			}
		}
	}

	if has, err := router.checkRouterReq(reqFilePath); err != nil {
		fmt.Printf("检查req失败: %v\n", err)
		return
	} else {
		if has {
			fmt.Printf("req已存在: %s\n", reqFilePath)
		} else {
			if err := filex.Append(reqFilePath, router.genRouterReq()); err != nil {
				fmt.Printf("创建req失败: %v\n", err)
				return
			}
		}
	}

	if has, err := router.checkRouterReply(replyFilePath); err != nil {
		fmt.Printf("检查reply失败: %v\n", err)
		return
	} else {
		if has {
			fmt.Printf("reply已存在: %s\n", replyFilePath)
		} else {
			if err := filex.Append(replyFilePath, router.genRouterReply()); err != nil {
				fmt.Printf("创建reply失败: %v\n", err)
				return
			}
		}
	}

	fmt.Println("Done!")
}
