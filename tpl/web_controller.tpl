package {{.Package}}

import (
	"{{.Project}}/logic"
	"{{.Project}}/types/req"

	"github.com/gin-gonic/gin"
	httpx "github.com/liuxiaobopro/gobox/http"
	respx "github.com/liuxiaobopro/gobox/resp"
)

type {{.Handle}}Handle struct {
	httpx.GinHandle
}

var {{.Controller}}Controller = &{{.Handle}}Handle{}

// IndexGet get请求
func (th *{{.Handle}}Handle) IndexGet(c *gin.Context) {
	var r req.{{.Controller}}GetReq
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnErr(c, respx.ParamErrT.ToPt())
		return
	}
	data, err := logic.{{.Controller}}logic.IndexGet(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

// IndexPost post请求
func (th *{{.Handle}}Handle) IndexPost(c *gin.Context) {
	var r req.{{.Controller}}PostReq
	if err := th.ShouldBindJSON(c, &r); err != nil {
		th.ReturnErr(c, respx.ParamErrT.ToPt())
		return
	}
	data, err := logic.{{.Controller}}logic.IndexPost(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}
