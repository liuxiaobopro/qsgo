package web

import (
	"fmt"

	filex "github.com/liuxiaobopro/gobox/file"
	stringx "github.com/liuxiaobopro/gobox/string"
)

func router(name string) {
	fmt.Printf("开始创建router: %s\n", name)
	if !stringx.Has(name, '/') {
		fmt.Printf("router格式错误,例如: demo/%s, 就是在demo.go添加%s方法", name, name)
		return
	}

	path := stringx.CutStartString(name, '/')
	path = path[:len(path)-1]
	handleFunc := stringx.CutEndString(name, '/')

	controllerFilePath := fmt.Sprintf("./controller/%s.go", path)
	logicFilePath := fmt.Sprintf("./logic/%s.go", path)
	reqFilePath := "./types/req/req.go"
	replyFilePath := "./types/reply/reply.go"

	router := &genRouter{
		Handle:      "demo",
		HandleFunc:  handleFunc,
		HandleLogic: "logic",
		Logic:       "demo",
		LogicFunc:   "Demo",
		Req:         "Demo" + handleFunc,
		Reply:       "Demo" + handleFunc,
	}

	if err := filex.Append(controllerFilePath, router.genRouterController()); err != nil {
		fmt.Printf("创建controller失败: %v\n", err)
		return
	}

	if err := filex.Append(logicFilePath, router.genRouterLogic()); err != nil {
		fmt.Printf("创建logic失败: %v\n", err)
		return
	}

	if err := filex.Append(reqFilePath, router.genRouterReq()); err != nil {
		fmt.Printf("创建req失败: %v\n", err)
		return
	}

	if err := filex.Append(replyFilePath, router.genRouterReply()); err != nil {
		fmt.Printf("创建reply失败: %v\n", err)
		return
	}

	fmt.Println("Done!")
}
