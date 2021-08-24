package record

import (
	"context"

	http2 "net/http"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Recorder interface {
	Record(context.Context, *http2.Request, interface{}, interface{}) error
	GetRawDataBeforeOperate(context.Context, interface{}) (interface{}, error)
}

type RecorderGenerator func(interface{}) (Recorder, error)

func NewRecordMiddleware(g RecorderGenerator) middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(*http.Transport); ok {
					var raw interface{}
					rd, err1 := g(req)
					if err1 != nil {
						return h(ctx, req)
					}
					if ht.Request().Method == "PUT" {
						// do something with the error
						raw, _ = rd.GetRawDataBeforeOperate(ctx, req)
					}
					defer func() {
						if err != nil {
							return
						}
						// do something with the error
						rd.Record(ctx, ht.Request(), req, raw)
					}()
				}
			}
			return h(ctx, req)
		}
	}
}
