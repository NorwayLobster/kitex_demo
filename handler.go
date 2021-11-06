package main

import (
	"context"
	"fmt"
	"log"
	"module_name_daytime/kitex_gen/api"
	"time"
)

// EchoImpl implements the last service interface defined in the IDL.
type ServiceNameDaytimeInThriftFileImpl struct{}

// Daytime implements the ServiceNameDaytimeInThriftFileImpl interface.
func (s *ServiceNameDaytimeInThriftFileImpl) Daytime(ctx context.Context, req *api.DaytimeRequest) (resp *api.DaytimeResponse, err error) {
	// TODO: Your code here...
	log.Printf("client:%s", req.ClientDaytime)
	daytime := time.Now()
	log.Printf("server:%v", daytime)
	resp = &api.DaytimeResponse{ServerDaytime: fmt.Sprintf("%v", daytime)}
	return
}
