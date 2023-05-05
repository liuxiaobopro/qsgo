package web

import (
	"fmt"
	"strings"

	stringx "github.com/liuxiaobopro/gobox/string"
	fmtp "github.com/liuxiaobopro/qsgo/log/fmt"
)

func Start(arg string) {
	if stringx.Count(arg, byte('=')) != 1 {
		fmt.Println("参数格式错误")
		return
	}
	argArr := strings.Split(arg, "=")
	fmtp.Printf("argArr: %v\n", argArr)
	switch argArr[0] {
	case "name":
		name(argArr[1])
	case "api":
		api(argArr[1])
	case "router":
		router(argArr[1])
	default:
		fmt.Println("参数未找到")
	}
}

func getLogicPath(name string) string {
	if stringx.Has(name, byte('/')) {
		s := stringx.CutStartString(name, '/')
		s = s[:len(s)-1]
		return "/" + s
	} else {
		return ""
	}
}

func getLogic(name string) string {
	if stringx.Has(name, byte('/')) {
		s := stringx.CutStartString(name, '/')
		s = stringx.ReplaceCharAfterSpecifiedCharLow(s, "/")
		return s + "Logic"
	} else {
		return "logic"
	}
}

func getPackage(name, d string) string {
	if !stringx.Has(name, byte('/')) {
		return d
	} else {
		arr := strings.Split(name, "/")
		return arr[len(arr)-2]
	}
}

func getHandle(name string) string {
	return stringx.ReplaceCharAfterSpecifiedCharLow(name, "/")
}

func getCL(name string) string {
	return stringx.ReplaceCharAfterSpecifiedCharUp(name, "/")
}

type genRouter struct {
	Handle      string
	HandleFunc  string
	HandleLogic string
	Logic       string
	LogicFunc   string
	Req         string
	Reply       string
}

func (th *genRouter) genRouterController() string {
	return `

func (th *` + th.Handle + `Handle) ` + th.HandleFunc + `(c *gin.Context) {
	var r req.` + th.Req + `Req
	if err := th.ShouldBind(c, &r); err != nil {  // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT.ToPt())
		return
	}
	data, err := ` + th.HandleLogic + `.` + th.LogicFunc + `logic.` + th.HandleFunc + `(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}
	`
}

func (th *genRouter) genRouterLogic() string {
	return `

func (th *` + th.Logic + `Logic) ` + th.LogicFunc + `(in *req.` + th.Req + `Req) (out *reply.` + th.Reply + `Reply, err *respx.Pt) {
	//TODO: write your logic here
	out = &reply.` + th.Reply + `Reply{}
	return
}
	`
}

func (th *genRouter) genRouterReq() string {
	return `

type ` + th.Req + `Req struct {}
	`
}

func (th *genRouter) genRouterReply() string {
	return `

type  ` + th.Reply + `Reply struct {}
	`
}
