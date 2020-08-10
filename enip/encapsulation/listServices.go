package encapsulation

import "github.com/MiguelValentine/goplc/enip/etype"

func (r *Request) ListServices(context uint64) []byte {
	pkg := &Encapsulation{}
	pkg.Command = CommandListServices
	pkg.SenderContext = context

	return pkg.Buffer()
}

func (r *Response) ListServices(context uint64, session etype.XUDINT, data []byte) []byte {
	pkg := &Encapsulation{}
	pkg.Command = CommandListServices
	pkg.SenderContext = context
	pkg.SessionHandle = session
	pkg.Data = data
	return pkg.Buffer()
}
