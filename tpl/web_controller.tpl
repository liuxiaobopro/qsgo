package {{.Package}}

import (
	"encoding/json"

	"{{.Project}}/define/types/req"
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
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	j, _ := json.Marshal(r)
	global.Logger.Infof(c, "{{.CL}}IndexReq: %s", j)
	data, err := {{.Logic}}.{{.CL}}logic.Index(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}