package {{.Package}}

import (
	{{.Logic}}"{{.Project}}/logic{{.LogicPath}}"
	"{{.Project}}/types/req"

	"github.com/gin-gonic/gin"
	httpx "github.com/liuxiaobopro/gobox/http"
	respx "github.com/liuxiaobopro/gobox/resp"
)

type {{.Handle}}Handle struct {
	httpx.GinHandle
}

var {{.CL}}Controller = &{{.Handle}}Handle{}

// Index Index
func (th *{{.Handle}}Handle) Index(c *gin.Context) {
	var r req.{{.CL}}IndexReq
	if err := th.ShouldBind(c, &r); err != nil { // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT.ToPt())
		return
	}
	data, err := {{.Logic}}.{{.CL}}logic.Index(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}