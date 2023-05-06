package {{.Package}}

import (
	"{{.Project}}/types/reply"
	"{{.Project}}/types/req"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type {{.Handle}}Logic struct{}

var {{.CL}}logic = &{{.Handle}}Logic{}

// Index Index
func (th *{{.Handle}}Logic) Index(in *req.{{.CL}}IndexReq) (*reply.{{.CL}}IndexReply, *respx.T) { // 最好保留一个func, 为了保留import
	//TODO: write your logic here
	out := &reply.{{.CL}}IndexReply{}
	return out, nil
}