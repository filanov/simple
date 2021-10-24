package hellower

import (
	"context"
	"simple/models"
	"simple/restapi"
	"simple/restapi/operations/hello"

	"github.com/go-openapi/runtime/middleware"
)

var _ restapi.HelloAPI = (*Hellower)(nil)

type Hellower struct{}

func (h *Hellower) GetHello(ctx context.Context, params hello.GetHelloParams) middleware.Responder {
	return hello.NewGetHelloOK().WithPayload(&models.HelloReply{Reply: "world"})
}
