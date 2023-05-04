package {{.Package}}

import (
	"{{.Project}}/types/reply"
	"{{.Project}}/types/req"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type {{.Handle}}Logic struct{}

var {{.CL}}logic = &{{.Handle}}Logic{}

// IndexGet get请求
func (th *{{.Handle}}Logic) IndexGet(in *req.{{.CL}}GetReq) (out *reply.{{.CL}}GetReply, err *respx.Pt) {
	//TODO: write your logic here
	out = &reply.{{.CL}}GetReply{
		Id:   in.Id,
		Name: "IndexGet",
	}
	return
}

// IndexPost post请求
func (th *{{.Handle}}Logic) IndexPost(in *req.{{.CL}}PostReq) (out *reply.{{.CL}}PostReply, err *respx.Pt) {
	//TODO: write your logic here
	out = &reply.{{.CL}}PostReply{
		Id:   in.Id,
		Name: "IndexPost",
	}
	return
}
