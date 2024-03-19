package web

import (
	"fmt"
	"path/filepath"
	"strings"

	filex "github.com/liuxiaobopro/gobox/file"
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
	fmtp.Printf("getLogic name: %s\n", name)
	if stringx.Has(name, byte('/')) {
		s := filepath.Dir(name)
		fmtp.Printf("getLogic s: %s\n", s)
		s = stringx.ReplaceCharAfterSpecifiedCharLow(s, "\\")
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
	Func     string // 方法名
	Handle   string // 处理器名
	CL       string // req reply
	Logic    string // logic
	LogicVar string // logic变量名
}

func (th *genRouter) genRouterController() string {
	return `

func (th *` + th.Handle + `Handle) ` + th.Func + `(c *gin.Context) {
	var r req.` + th.CL + `Req
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnStatusOKErr(c, replyx.ParamErrT)
		return
	}
	j, _ := json.Marshal(r)
	global.Logger.Debugf(c, "` + th.CL + `Req2: %s", j)
	c.Set(define.ReqText, string(j))
	data, err := ` + th.Logic + `.` + th.LogicVar + `logic.` + th.Func + `(c, &r)
	if err != nil {
		th.ReturnStatusOKErr(c, err)
		return
	}
	if m, err := define.DefaultResStyle(data); err != nil {
		th.ReturnStatusOKErr(c, replyx.InternalErrT)
	} else {
		th.RetuenOk(c, m)
	}
}
	`
}

func (th *genRouter) checkRouterController(filePath string) (bool, error) {
	s := `func (th *` + th.Handle + `Handle) ` + th.Func + ``
	fmtp.Printf("checkRouterController str: %v\n", s)
	return filex.Has(filePath, s)
}

func (th *genRouter) genRouterLogic() string {
	return `

func (th *` + th.Handle + `Logic) ` + th.Func + `(c *gin.Context, in *req.` + th.CL + `Req) (*reply.` + th.CL + `Reply, *replyx.T) {
	//TODO: write your logic here
	var(
		out = &reply.` + th.CL + `Reply{}
	)
	return out, nil
}
	`
}

func (th *genRouter) checkRouterLogic(filePath string) (bool, error) {
	s := `func (th *` + th.Handle + `Logic) ` + th.Func + ``
	fmtp.Printf("checkRouterLogic str: %v\n", s)
	return filex.Has(filePath, s)
}

func (th *genRouter) genRouterReq() string {
	return `

type ` + th.CL + `Req struct {}
	`
}

func (th *genRouter) checkRouterReq(filePath string) (bool, error) {
	s := `type ` + th.CL + `Req struct`
	fmtp.Printf("checkRouterReq str: %v\n", s)
	return filex.Has(filePath, s)
}

func (th *genRouter) genRouterReply() string {
	return `

type ` + th.CL + `Reply struct {}
	`
}

func (th *genRouter) checkRouterReply(filePath string) (bool, error) {
	s := `type ` + th.CL + `Reply struct`
	fmtp.Printf("checkRouterReply str: %v\n", s)
	return filex.Has(filePath, s)
}
