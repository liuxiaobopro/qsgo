package web

import (
	"fmt"
	"strings"

	stringx "github.com/liuxiaobopro/gobox/string"
)

var ()

func Start(arg string) {
	if stringx.Count(arg, byte('=')) != 1 {
		fmt.Println("参数格式错误")
		return
	}
	argArr := strings.Split(arg, "=")
	switch argArr[0] {
	case "name":
		name(argArr[1])
	case "api":
		api(argArr[1])
	default:
		fmt.Println("参数未找到")
	}
}
