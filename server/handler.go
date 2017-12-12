package server

import (
	"../api"
	"context"
	"fmt"
	"time"
)

var conn = 0

type GatewayHandler struct {
}

func (h *GatewayHandler) Start(ctx context.Context) (r *api.Result_, err error) {
	time.Sleep(time.Second * 5)

	ret := &api.Result_{Ret: api.ResultCode_Okay, Conn: 88}
	fmt.Printf("start gateway %d\n", conn)
	return ret, nil
}
