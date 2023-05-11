package {{.Package}}

import (
	"{{.Project}}/define/types/reply"
	"{{.Project}}/define/types/req"

	replyx "github.com/liuxiaobopro/gobox/reply"
	"github.com/gin-gonic/gin"
)

type {{.Handle}}Logic struct{}

var {{.CL}}logic = &{{.Handle}}Logic{}

// Index Index
func (th *{{.Handle}}Logic) Index(c *gin.Context, in *req.{{.CL}}IndexReq) (*reply.{{.CL}}IndexReply, *replyx.T) { // 最好保留一个func, 为了保留import
	//TODO: write your logic here
	out := &reply.{{.CL}}IndexReply{}
	return out, nil
}