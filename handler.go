package main

import (
	"context"
	"module_name_daytime/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Daytime implements the EchoImpl interface.
func (s *EchoImpl) Daytime(ctx context.Context, req *api.DaytimeRequest) (resp *api.DaytimeResponse, err error) {
	// TODO: Your code here...
	return
}

// Daytime implements the ServiceNameDaytimeIntThriftFileImpl interface.
func (s *ServiceNameDaytimeIntThriftFileImpl) Daytime(ctx context.Context, req *api.DaytimeRequest) (resp *api.DaytimeResponse, err error) {
	// TODO: Your code here...
	return
}

// Daytime implements the ServiceNameDaytimeInThriftFileImpl interface.
func (s *ServiceNameDaytimeInThriftFileImpl) Daytime(ctx context.Context, req *api.DaytimeRequest) (resp *api.DaytimeResponse, err error) {
	// TODO: Your code here...
	return
}
