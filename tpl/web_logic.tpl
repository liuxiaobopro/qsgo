package {{.Package}}

import (
	"{{.Project}}/types/reply"
	"{{.Project}}/types/req"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type {{.Handle}}Logic struct{}

var {{.CL}}logic = &{{.Handle}}Logic{}

// Index Index
func (th *{{.Handle}}Logic) Index(in *req.{{.CL}}IndexReq) (out *reply.{{.CL}}IndexReply, err *respx.Pt) {
	//TODO: write your logic here
	out = &reply.{{.CL}}IndexReply{}
	return
}