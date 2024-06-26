package {{.Package}}

import (
	"encoding/json"

	"{{.Project}}/define/types/req"
	"{{.Project}}/define"
	"{{.Project}}/global"
	{{.Logic}}"{{.Project}}/logic{{.LogicPath}}"

	"github.com/gin-gonic/gin"
	httpx "github.com/liuxiaobopro/gobox/http"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

type {{.Handle}}Handle struct {
	httpx.GinHandle
}

var {{.CL}}Controller = &{{.Handle}}Handle{}

// Index Index
func (th *{{.Handle}}Handle) Index(c *gin.Context) { // 最好保留一个func, 为了保留import
	var r req.{{.CL}}IndexReq
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnStatusOKErr(c, replyx.ParamErrT)
		return
	}
	j, _ := json.Marshal(r)
	global.Logger.Debugf(c, "{{.CL}}IndexReq: %s", j)
	c.Set(define.ReqText, string(j))
	data, err := {{.Logic}}.{{.CL}}logic.Index(c, &r)
	if err != nil {
		th.ReturnStatusOKErr(c, err)
		return
	}
	if m, err := define.DefaultResStyle(data); err != nil {
		th.ReturnStatusOKErr(c, replyx.InternalErrT)
	} else {
		th.ReturnOk(c, m)
	}
}